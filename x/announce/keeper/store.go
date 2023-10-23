package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func (k Keeper) getAnnouncementsStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.AnnouncedStorageLocations)
}

// getAnnouncedValidators unmarshal announced validators from storage
func (k Keeper) getAnnouncedValidators(ctx sdk.Context) *types.GetAnnouncedValidatorsResponse {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AnnouncedStorageLocations)
	defer iterator.Close()
	validators := []string{}

	for ; iterator.Valid(); iterator.Next() {
		validator := iterator.Key()
		validator = bytes.TrimPrefix(validator, types.AnnouncedStorageLocations)
		validators = append(validators, hexutil.Encode(validator))
	}

	announcedValidators := &types.GetAnnouncedValidatorsResponse{
		Validator: validators,
	}
	return announcedValidators
}

// getAnnouncements unmarshal announcements for the given validator from storage
func (k Keeper) getAnnouncements(ctx sdk.Context, validator []byte) (*types.StoredAnnouncements, error) {
	store := k.getAnnouncementsStore(ctx)
	announceBytes := store.Get(validator)
	announcements := &types.StoredAnnouncements{}

	if announceBytes == nil {
		return announcements, nil
	}

	err := announcements.Unmarshal(announceBytes)
	return announcements, err
}

// setAnnouncements store the announcements for the given validator
func (k Keeper) setAnnouncements(ctx sdk.Context, validator []byte, announcements *types.StoredAnnouncements) error {
	store := k.getAnnouncementsStore(ctx)
	announcementsBytes, err := announcements.Marshal()
	if err != nil {
		return err
	}
	store.Set(validator, announcementsBytes)
	return nil
}
