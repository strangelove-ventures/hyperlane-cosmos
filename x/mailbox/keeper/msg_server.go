package keeper

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"

	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.MsgServer = (*Keeper)(nil)

const MAX_MESSAGE_BODY_BYTES = 2_000

type ContractProcessMsg struct {
	Origin uint32 `json:"origin"`
	Sender string `json:"sender"`
	Msg    string `json:"msg"`
}

// NewMsgServerImpl return an implementation of the mailbox MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// Dispatch defines a rpc handler method for MsgDispatch
func (k Keeper) Dispatch(goCtx context.Context, msg *types.MsgDispatch) (*types.MsgDispatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: NewMessage
	var message []byte

	// TODO: Make sure this is the right version
	version := make([]byte, 1)
	message = append(message, version...)

	// Nonce is the tree count.
	nonce := uint32(k.Tree.Count())
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, nonce)
	message = append(message, nonceBytes...)

	// Local Domain is set on NewKeeper
	origin := uint32(k.domain)
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	message = append(message, originBytes...)

	// Get the Sender address
	// Since this is a cosmos chain, sender will be a bech32 address
	sender := sdk.MustAccAddressFromBech32(msg.Sender).Bytes()
	for len(sender) < (ismtypes.DESTINATION_OFFSET - ismtypes.SENDER_OFFSET) {
		padding := make([]byte, 1)
		sender = append(padding, sender...)
	}
	message = append(message, sender...)

	// Get the Destination Domain
	destination := msg.DestinationDomain
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, destination)
	message = append(message, destinationBytes...)

	// Get the Recipient address
	// Since the recipient could be any destination change, the address must be in hex, non-bech32 format
	recipient := hexutil.MustDecode(msg.RecipientAddress)
	for len(recipient) < (ismtypes.BODY_OFFSET - ismtypes.RECIPIENT_OFFSET) {
		padding := make([]byte, 1)
		recipient = append(padding, recipient...)
	}
	message = append(message, recipient...)

	// Get the Message Body
	// messageBytes := []byte(msg.MessageBody)
	messageBytes := hexutil.MustDecode(msg.MessageBody)
	if len(messageBytes) > MAX_MESSAGE_BODY_BYTES {
		return nil, types.ErrMsgTooLong
	}
	message = append(message, messageBytes...)

	// Get the message ID
	id := ismtypes.Id(message)

	// Insert the message id into the tree
	err := k.Tree.Insert(id)
	if err != nil {
		return nil, err
	}
	// Store that the leaf
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MailboxIMTKey(k.Tree.Count()-1), id)

	// Emit the events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDispatch,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyDestinationDomain, strconv.FormatUint(uint64(msg.DestinationDomain), 10)),
			sdk.NewAttribute(types.AttributeKeyRecipientAddress, msg.RecipientAddress),
			sdk.NewAttribute(types.AttributeKeyMessage, msg.MessageBody),
		),
		sdk.NewEvent(
			types.EventTypeDispatchId,
			sdk.NewAttribute(types.AttributeKeyID, string(hexutil.Encode(id))),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgDispatchResponse{
		MessageId: hexutil.Encode(id),
	}, nil
}

// Process defines a rpc handler method for MsgProcess
func (k Keeper) Process(goCtx context.Context, msg *types.MsgProcess) (*types.MsgProcessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	messageBytes := hexutil.MustDecode(msg.Message)
	id, err := k.VerifyMessage(messageBytes)
	if err != nil {
		return nil, err
	}

	metadataBytes := hexutil.MustDecode(msg.Metadata)
	// Verify message signatures
	if !k.ismKeeper.Verify(metadataBytes, messageBytes) {
		return nil, types.ErrMsgVerificationFailed
	}

	// Parse the recipient and get the contract address
	recipientBytes := ismtypes.Recipient(messageBytes)
	recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), recipientBytes)
	contractAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "contract")
	}

	// Parse origin, sender, and body for the contract msg
	origin := ismtypes.Origin(messageBytes)
	senderBytes := ismtypes.Sender(messageBytes)
	senderHex := hexutil.Encode(senderBytes)
	body := ismtypes.Body(messageBytes)
	contractMsg := ContractProcessMsg{
		Origin: origin,
		Sender: senderHex,
		Msg:    string(body),
	}
	encodedMsg, err := json.Marshal(contractMsg)
	if err != nil {
		return nil, err
	}

	// Call the recipient contract
	_, err = k.pcwKeeper.Execute(ctx, contractAddr, k.mailboxAddr, encodedMsg, sdk.NewCoins())
	if err != nil {
		return nil, err
	}

	// Store that the message was delivered
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MailboxDeliveredKey(id), []byte{1})
	k.Delivered[id] = true

	// Emit the events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
		sdk.NewEvent(
			types.EventTypeProcessId,
			sdk.NewAttribute(types.AttributeKeyID, id),
		),
		sdk.NewEvent(
			types.EventTypeProcess,
			sdk.NewAttribute(types.AttributeKeyOrigin, strconv.FormatUint(uint64(origin), 10)),
			sdk.NewAttribute(types.AttributeKeySender, senderHex),
			sdk.NewAttribute(types.AttributeKeyRecipientAddress, recipient),
			sdk.NewAttribute(types.AttributeKeyMessage, string(body)),
		),
	})

	return &types.MsgProcessResponse{}, nil
}
