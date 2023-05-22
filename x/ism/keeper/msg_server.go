package keeper

import (
	"context"
	"encoding/json"

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

	ismMap := map[uint32][]byte{}
	for _, originIsm := range msg.Isms {
		ism, err := k.cdc.Marshal(originIsm.Ism)
		if err != nil {
			return &types.MsgSetDefaultIsmResponse{}, err
		}
		ismMap[originIsm.Origin] = ism
	}

	eventValue, err := json.Marshal(msg.Isms)
	if err != nil {
		return &types.MsgSetDefaultIsmResponse{}, err
	}

	store := ctx.KVStore(k.storeKey)
	for _, originIsm := range msg.Isms {
		k.defaultIsm[originIsm.Origin] = *originIsm.Ism
		store.Set(types.OriginKey(originIsm.Origin), ismMap[originIsm.Origin])
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetDefaultIsm,
			sdk.NewAttribute(types.AttributeKeySetDefaultIsm, string(eventValue)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgSetDefaultIsmResponse{}, nil
}
