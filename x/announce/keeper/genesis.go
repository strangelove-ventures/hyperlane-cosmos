package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

// InitGenesis initializes the hyperlane announce module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	for _, announcement := range gs.Announcements {
		err := k.setAnnouncements(ctx, announcement.Validator, announcement.Announcements)
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
		validator = bytes.TrimPrefix(validator, types.AnnouncedStorageLocations)

		storedAnnouncementsBytes := iterator.Value()
		storedAnnouncements := &types.StoredAnnouncements{}
		err := storedAnnouncements.Unmarshal(storedAnnouncementsBytes)
		if err != nil {
			return nil, err
		}

		currAnnouncement := &types.GenesisAnnouncement{
			Announcements: storedAnnouncements,
			Validator:     validator,
		}
		announcements = append(announcements, currAnnouncement)
	}
	return announcements, nil
}
