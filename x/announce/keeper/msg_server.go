package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the announce MsgServer interface for the provided keeper
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
	} else {
		// Check for replays since there were existing announcements for this validator
		for _, announcement := range storedAnnouncements.Announcement {
			if announcement.StorageLocation == msg.StorageLocation {
				return nil, types.ErrReplayAnnouncement
			}
		}
	}

	origin := k.mailboxKeeper.GetDomain(ctx)
	mailboxAddr := k.mailboxKeeper.GetMailboxAddress()

	announcementDigest, err := types.GetAnnouncementDigest(origin, mailboxAddr, msg.StorageLocation)
	if err != nil {
		return nil, err
	}
	err = types.VerifyAnnouncementDigest(announcementDigest, msg.Signature, msg.Validator)
	if err != nil {
		return nil, err
	}

	storedAnnouncements.Announcement = append(storedAnnouncements.Announcement, &types.StoredAnnouncement{
		StorageLocation: msg.StorageLocation,
		Signature:       msg.Signature,
	})

	err = k.setAnnouncedValidators(ctx, msg.Validator)
	if err != nil {
		return nil, types.ErrMarshalAnnouncedValidators
	}
	err = k.setAnnouncements(ctx, msg.Validator, storedAnnouncements)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAnnounce,
			sdk.NewAttribute(types.AttributeStorageLocation, msg.StorageLocation),
			sdk.NewAttribute(types.AttributeValidatorAddress, hex.EncodeToString(msg.Validator)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgAnnouncementResponse{}, nil
}
