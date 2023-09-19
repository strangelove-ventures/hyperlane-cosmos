package bindings

import (
	"encoding/json"

	sdkerrors "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	bindingstypes "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/bindings/types"
	mailboxkeeper "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/keeper"
	mailboxtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(mailbox *mailboxkeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped: old,
			mailbox: mailbox,
		}
	}
}

type CustomMessenger struct {
	wrapped wasmkeeper.Messenger
	mailbox *mailboxkeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle the happy path, leave everything else for the wrapped version
		var contractMsg bindingstypes.MailboxMsgType
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "mailbox msg")
		}

		if contractMsg.MsgDispatch != nil {
			return m.msgDispatch(ctx, contractAddr, contractMsg.MsgDispatch)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

// msgDispatch calls into msg server's Dispatch
func (m *CustomMessenger) msgDispatch(ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindingstypes.MsgDispatch) ([]sdk.Event, [][]byte, error) {
	_, err := MsgDispatch(m.mailbox, ctx, contractAddr, msg)
	if err != nil {
		return nil, nil, sdkerrors.Wrap(err, "mailbox-specific msg dispatch")
	}
	// TODO: double check how this is all encoded to the contract
	return nil, nil, nil
}

func MsgDispatch(k *mailboxkeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindingstypes.MsgDispatch) ([]byte, error) {
	if msg == nil {
		return nil, wasmvmtypes.InvalidRequest{Err: "msgdispatch null msg"}
	}

	msgServer := mailboxkeeper.NewMsgServerImpl(k)

	msgMsgDispatch := mailboxtypes.NewMsgDispatch(msg.Sender, msg.DestinationDomain, msg.RecipientAddress, msg.MessageBody)

	if err := msgMsgDispatch.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrap(err, "failed validating MsgDispatch")
	}

	// Dispatch msg
	_, err := msgServer.Dispatch(
		sdk.WrapSDKContext(ctx),
		msgMsgDispatch,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "msg server dispatch msg")
	}

	return nil, nil
}
