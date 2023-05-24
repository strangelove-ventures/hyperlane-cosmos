package keeper

import (
	"context"
	"encoding/binary"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	ism "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.MsgServer = (*Keeper)(nil)

const MAX_MESSAGE_BODY_BYTES = 2_000

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
	// TODO: How to handle various address types?
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	message = append(message, sender.Bytes()...)

	// Get the Destination Domain
	destination := msg.DestinationDomain
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, destination)
	message = append(message, destinationBytes...)

	// Get the Recipient address
	// TODO: How to handle various address types?
	recipient := sdk.MustAccAddressFromBech32(msg.RecipientAddress)
	message = append(message, recipient.Bytes()...)

	// Get the Message Body
	messageBytes := hexutil.MustDecode(msg.MessageBody)
	if len(messageBytes) > MAX_MESSAGE_BODY_BYTES {
		return nil, types.ErrMsgTooLong
	}
	message = append(message, messageBytes...)

	// Insert the message into the tree
	err := k.tree.Insert(message)
	if err != nil {
		return nil, err
	}

	// Get the message ID
	id := ism.Id(message)

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
			types.EventTypeDispatch,
			sdk.NewAttribute(types.AttributeKeyID, string(hexutil.Encode(id))),
		),
	})

	return &types.MsgDispatchResponse{
		MessageId: hexutil.Encode(id),
	}, nil
}

// Process defines a rpc handler method for MsgProcess
func (k Keeper) Process(goCtx context.Context, msg *types.MsgProcess) (*types.MsgProcessResponse, error) {
	panic("Implement me")
}
