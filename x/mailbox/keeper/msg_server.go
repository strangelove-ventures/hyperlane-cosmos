package keeper

import (
	"context"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// Dispatch defines a rpc handler method for MsgDispatch
func (k Keeper) Dispatch(goCtx context.Context, msg *types.MsgDispatch) (*types.MsgDispatchResponse, error) {
	panic("Implement me")
}

// Process defines a rpc handler method for MsgProcess
func (k Keeper) Process(goCtx context.Context, msg *types.MsgProcess) (*types.MsgProcessResponse, error) {
	panic("Implement me")
}
