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

// StoreCode defines a rpc handler method for MsgStoreCode
func (k Keeper) SetDefaultIsm(goCtx context.Context, msg *types.MsgSetDefaultIsm) (*types.MsgSetDefaultIsmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Signer {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority: expected %s, got %s", k.authority, msg.Signer)
	}

	ismBzMap := map[uint32][]byte{}
	ismMap := map[uint32]types.AbstractIsm{}
	events := sdk.Events{}
	for _, originIsm := range msg.Isms {
		ism, err := types.UnpackAbstractIsm(originIsm.AbstractIsm)
		if err != nil {
			return &types.MsgSetDefaultIsmResponse{}, err
		}
		ismMap[originIsm.Origin] = ism

		ismBz, err := k.cdc.MarshalInterface(ism)
		if err != nil {
			return &types.MsgSetDefaultIsmResponse{}, err
		}
		ismBzMap[originIsm.Origin] = ismBz
		events.AppendEvent(ism.Event(originIsm.Origin))
	}

	store := ctx.KVStore(k.storeKey)
	for _, originIsm := range msg.Isms {
		k.defaultIsm[originIsm.Origin] = ismMap[originIsm.Origin]
		store.Set(types.OriginKey(originIsm.Origin), ismBzMap[originIsm.Origin])
	}

	events.AppendEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	ctx.EventManager().EmitEvents(events)

	return &types.MsgSetDefaultIsmResponse{}, nil
}
