package keeper

import (
	"context"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the ism MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// PayForGas defines a rpc handler method for MsgPayForGas
func (k Keeper) PayForGas(goCtx context.Context, msg *types.MsgPayForGas) (*types.MsgPayForGasResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgPayForGasResponse{}, nil
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
