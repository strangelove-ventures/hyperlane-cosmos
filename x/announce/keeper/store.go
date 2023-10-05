package keeper

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func (k Keeper) getAnnouncementsStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.AnnouncedStorageLocations)
}

// getAnnouncedValidators unmarshal announced validators from storage
func (k Keeper) getAnnouncedValidators(ctx sdk.Context) (*types.GetAnnouncedValidatorsResponse, error) {
	store := ctx.KVStore(k.storeKey)
	announcedValidatorBytes := store.Get(types.AnnouncedValidators)
	announcedValidators := &types.GetAnnouncedValidatorsResponse{}

	if announcedValidatorBytes == nil {
		return nil, errors.New("No announced validators")
	}

	err := announcedValidators.Unmarshal(announcedValidatorBytes)
	return announcedValidators, err
}

// setAnnouncedValidators store an announced validator
func (k Keeper) setAnnouncedValidators(ctx sdk.Context, validator []byte) (err error) {
	var announcedValidators *types.GetAnnouncedValidatorsResponse
	announcedValidators, err = k.getAnnouncedValidators(ctx)
	if err != nil {
		announcedValidators = &types.GetAnnouncedValidatorsResponse{Validator: [][]byte{}}
	}
	announcedValidators.Validator = append(announcedValidators.Validator, validator)

	announcedValidatorsBytes, err := announcedValidators.Marshal()
	if err != nil {
		return types.ErrMarshalAnnouncedValidators
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AnnouncedValidators, announcedValidatorsBytes)
	return nil
}

// getAnnouncements unmarshal announcements for the given validator from storage
func (k Keeper) getAnnouncements(ctx sdk.Context, validator []byte) (*types.StoredAnnouncements, error) {
	store := k.getAnnouncementsStore(ctx)
	announceBytes := store.Get(validator)
	announcements := &types.StoredAnnouncements{}

	if announceBytes == nil {
		return nil, errors.New("No announcements stored for validator")
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
