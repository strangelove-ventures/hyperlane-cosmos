package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the ism MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// SetDefaultIsm defines a rpc handler method for MsgSetDefaultIsm
func (k Keeper) SetDefaultIsm(goCtx context.Context, msg *types.MsgSetDefaultIsm) (*types.MsgSetDefaultIsmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Signer {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority: expected %s, got %s", k.authority, msg.Signer)
	}

	events := sdk.Events{}
	for _, originIsm := range msg.Isms {
		ism, err := types.UnpackAbstractIsm(originIsm.AbstractIsm)
		if err != nil {
			return &types.MsgSetDefaultIsmResponse{}, err
		}
		err = k.storeDefaultIsm(ctx, originIsm.Origin, ism)
		if err != nil {
			return &types.MsgSetDefaultIsmResponse{}, err
		}

		events = events.AppendEvent(ism.DefaultIsmEvent(originIsm.Origin))
	}

	events = events.AppendEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	ctx.EventManager().EmitEvents(events)

	return &types.MsgSetDefaultIsmResponse{}, nil
}

// CreateIsm defines a rpc handler method for MsgCreateIsm
func (k Keeper) CreateIsm(goCtx context.Context, msg *types.MsgCreateIsm) (*types.MsgCreateIsmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ismId, err := k.getNextCustomIsmIndex(ctx)
	if err != nil {
		return nil, err
	}

	ism, err := types.UnpackAbstractIsm(msg.Ism)
	if err != nil {
		return nil, err
	}

	err = k.storeCustomIsm(ctx, ismId, ism)
	if err != nil {
		return nil, err
	}

	events := sdk.Events{}
    events = events.AppendEvent(ism.CustomIsmEvent(ismId))

	events = events.AppendEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	ctx.EventManager().EmitEvents(events)

	return &types.MsgCreateIsmResponse{
		IsmId: ismId,
	}, nil
}
