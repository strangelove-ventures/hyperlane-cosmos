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
func (k Keeper) Announcement(goCtx context.Context, msg *types.MsgAnnouncement) (resp *types.MsgAnnouncementResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var storedAnnouncements *types.StoredAnnouncements
	storedAnnouncements, err = k.getAnnouncements(ctx, msg.Validator)
	if err != nil {
		storedAnnouncements = &types.StoredAnnouncements{
			Announcement: []*types.StoredAnnouncement{},
		}
	}

	// TODO: verify properly before storing ...

	storedAnnouncements.Announcement = append(storedAnnouncements.Announcement, &types.StoredAnnouncement{
		StorageLocation: msg.StorageLocation,
		Signature:       msg.Signature,
	})
	err = k.setAnnouncements(ctx, msg.Validator, storedAnnouncements)
	if err != nil {
		return nil, types.ErrMarshalAnnouncement
	}

	return &types.MsgAnnouncementResponse{}, nil
}
