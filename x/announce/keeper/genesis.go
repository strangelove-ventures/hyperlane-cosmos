package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

// InitGenesis initializes the hyperlane announce module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	var storedAnnouncements *types.StoredAnnouncements
	var err error

	for _, announcement := range gs.Announcements {
		storedAnnouncements, err = k.getAnnouncements(ctx, announcement.Validator)
		if err != nil {
			storedAnnouncements = &types.StoredAnnouncements{
				Announcement: []*types.StoredAnnouncement{},
			}
		} else {
			// Check for replays since there were existing announcements for this validator
			for _, existingAnnouncement := range storedAnnouncements.Announcement {
				if existingAnnouncement.StorageLocation == announcement.Announcement.StorageLocation {
					return types.ErrReplayAnnouncement
				}
			}
		}

		storedAnnouncements.Announcement = append(storedAnnouncements.Announcement, &types.StoredAnnouncement{
			StorageLocation: announcement.Announcement.StorageLocation,
		})

		err = k.setAnnouncements(ctx, announcement.Validator, storedAnnouncements)
		if err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis returns the hyperlane announce module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	announcements, err := ExportAnnouncements(ctx.KVStore(k.storeKey))
	if err != nil {
		panic("Exporting genesis state failed - exporting announcements failed")
	}
	return types.GenesisState{
		Announcements: announcements,
	}
}

// getAnnouncedValidators unmarshal announced validators from storage
func ExportAnnouncements(store sdk.KVStore) ([]*types.GenesisAnnouncement, error) {
	iterator := sdk.KVStorePrefixIterator(store, types.AnnouncedStorageLocations)
	defer iterator.Close()
	announcements := []*types.GenesisAnnouncement{}

	for ; iterator.Valid(); iterator.Next() {
		validator := iterator.Key()
		storedAnnouncementsBytes := iterator.Value()
		storedAnnouncements := &types.StoredAnnouncements{}
		err := storedAnnouncements.Unmarshal(storedAnnouncementsBytes)
		if err != nil {
			return nil, err
		}
		for _, curr := range storedAnnouncements.Announcement {
			currAnnouncement := &types.GenesisAnnouncement{
				Announcement: curr,
				Validator:    validator,
			}
			announcements = append(announcements, currAnnouncement)
		}
	}
	return announcements, nil
}
