package keeper

import (
	"context"
	"encoding/binary"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common/hexutil"

	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.MsgServer = (*Keeper)(nil)

const MAX_MESSAGE_BODY_BYTES = 2_000

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
	nonce := uint32(k.tree.Count())
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
	//messageBytes := []byte(msg.MessageBody)
	messageBytes := hexutil.MustDecode(msg.MessageBody)
	if len(messageBytes) > MAX_MESSAGE_BODY_BYTES {
		return nil, types.ErrMsgTooLong
	}
	message = append(message, messageBytes...)

	// Get the message ID
	id := ismtypes.Id(message)

	// Insert the message id into the tree
	err := k.tree.Insert(id)
	if err != nil {
		return nil, err
	}

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
	if ismtypes.Version(messageBytes) != k.version {
		return nil, types.ErrMsgInvalidVersion
	}

	if ismtypes.Destination(messageBytes) != k.domain {
		return nil, types.ErrMsgInvalidDomain
	}

	idBytes := ismtypes.Id(messageBytes)
	id := hexutil.Encode(idBytes)

	// Let's make sure we've not already delivered the message
	// TODO: Load from Store
	val, ok := k.delivered[id]
	if ok && val {
		return nil, types.ErrMsgDelivered
	}

	// TODO: Store
	k.delivered[id] = true

	// TODO: GetRecipientISM
	i := k.getRecipientISM()

	metadataBytes := hexutil.MustDecode(msg.Metadata)
	if !i.Verify(metadataBytes, messageBytes) {
		return nil, types.ErrMsgVerificationFailed
	}

	origin := ismtypes.Origin(messageBytes)
	sender := ismtypes.Recipient(messageBytes)
	recipient := ismtypes.Recipient(messageBytes)
	body := ismtypes.Body(messageBytes)

	// TODO: Do we need to do anything with the data?
	_, err := k.HandleMessage(goCtx, origin, sender, recipient, string(body))
	if err != nil {
		return nil, types.ErrMsgHandling
	}

	// Emit the events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeProcess,
			sdk.NewAttribute(types.AttributeKeyOrigin, strconv.FormatUint(uint64(origin), 10)),
			sdk.NewAttribute(types.AttributeKeySender, sender),
			sdk.NewAttribute(types.AttributeKeyRecipientAddress, recipient),
			sdk.NewAttribute(types.AttributeKeyMessage, string(body)),
		),
		sdk.NewEvent(
			types.EventTypeProcessId,
			sdk.NewAttribute(types.AttributeKeyID, id),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})
	return &types.MsgProcessResponse{}, nil
}
