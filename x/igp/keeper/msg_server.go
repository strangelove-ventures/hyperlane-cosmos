package keeper

import (
	"context"
	"encoding/binary"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the ism MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// PayForGas defines a rpc handler method for MsgPayForGas
func (k Keeper) PayForGas(goCtx context.Context, msg *types.MsgPayForGas) (*types.MsgPayForGasResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	// events := sdk.Events{}
	// store := ctx.KVStore(k.storeKey)

	// store.Set(types.OriginKey(originIsm.Origin), ismBzMap[originIsm.Origin])

	return &types.MsgPayForGasResponse{}, nil
}

func (k Keeper) SetDestinationGasOverhead(goCtx context.Context, msg *types.MsgSetDestinationGasOverhead) (*types.MsgSetDestinationGasOverheadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	key := types.GasOverheadKey(msg.DestinationDomain)
	gasOverhead := make([]byte, 8)
	binary.LittleEndian.PutUint64(gasOverhead, msg.GasOverhead)
	store.Set(key, gasOverhead)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetGasOverhead,
			sdk.NewAttribute(types.AttributeDestination, strconv.FormatUint(uint64(msg.DestinationDomain), 10)),
			sdk.NewAttribute(types.AttributeOverheadAmount, strconv.FormatUint(uint64(msg.GasOverhead), 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})
	return &types.MsgSetDestinationGasOverheadResponse{}, nil
}

// Claim defines a rpc handler method for MsgClaims
func (k Keeper) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgClaimResponse{}, nil
}

// SetGasOracles defines a rpc handler method for MsgSetGasOracles
func (k Keeper) SetGasOracles(goCtx context.Context, msg *types.MsgSetGasOracles) (*types.MsgSetGasOraclesResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgSetGasOraclesResponse{}, nil
}

// SetGasOracles defines a rpc handler method for MsgSetGasOracles
func (k Keeper) SetBeneficiary(goCtx context.Context, msg *types.MsgSetBeneficiary) (*types.MsgSetBeneficiaryResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgSetBeneficiaryResponse{}, nil
}
