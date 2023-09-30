package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the igp MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// Announcement declares a storage location for a hyperlane validator
func (k Keeper) Announcement(goCtx context.Context, msg *types.MsgAnnouncement) (*types.MsgAnnouncementResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)
	return &types.MsgAnnouncementResponse{}, nil
}
