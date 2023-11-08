package keeper

import (
	"bytes"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

// getDefaultIsm returns the default ISM
func (k Keeper) getDefaultIsm(ctx sdk.Context, origin uint32) (types.AbstractIsm, error) {
	store := ctx.KVStore(k.storeKey)
	ismBz := store.Get(types.DefaultIsmKey(origin))

	var ism types.AbstractIsm
	err := k.cdc.UnmarshalInterface(ismBz, &ism)
	if err != nil {
		return nil, err
	}

	return ism, nil
}

// getAllDefaultIsms returns the default ISMs
func (k Keeper) getAllDefaultIsms(ctx sdk.Context) ([]*types.DefaultIsm, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyOriginsDefaultIsm))
	defer iterator.Close()

	var defaultIsms []*types.DefaultIsm
	for ; iterator.Valid(); iterator.Next() {
		originBz := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyOriginsDefaultIsm)))
		origin64, err := strconv.ParseUint(string(originBz), 10, 32)
		if err != nil {
			return nil, err
		}
		
		var ism types.AbstractIsm
		err = k.cdc.UnmarshalInterface(iterator.Value(), &ism)
		if err != nil {
			return nil, err
		}
		
		ismAny, err := types.PackAbstractIsm(ism)
		if err != nil {
			return nil, err
		}

		defaultIsms = append(defaultIsms, &types.DefaultIsm{
			Origin: uint32(origin64),
			AbstractIsm: ismAny,
		})
	}

	return defaultIsms, nil
}

// setDefaultIsm stores the default ISM
func (k Keeper) storeDefaultIsm(ctx sdk.Context, origin uint32, ism types.AbstractIsm) error {
	store := ctx.KVStore(k.storeKey)

	ismBz, err := k.cdc.MarshalInterface(ism)
	if err != nil {
		return err
	}
	store.Set(types.DefaultIsmKey(origin), ismBz)

	return nil
}

// getCustomIsm returns the custom ISM
func (k Keeper) getCustomIsm(ctx sdk.Context, ismId uint32) (types.AbstractIsm, error) {
	store := ctx.KVStore(k.storeKey)
	ismBz := store.Get(types.CustomIsmKey(ismId))

	var ism types.AbstractIsm
	err := k.cdc.UnmarshalInterface(ismBz, &ism)
	if err != nil {
		return nil, err
	}

	return ism, nil
}

// getAllCustomIsms returns the custom ISMs
func (k Keeper) getAllCustomIsms(ctx sdk.Context) ([]*types.CustomIsm, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyCustomIsm))
	defer iterator.Close()

	var customIsms []*types.CustomIsm
	for ; iterator.Valid(); iterator.Next() {
		indexBz := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyCustomIsm)))
		index64, err := strconv.ParseUint(string(indexBz), 10, 32)
		if err != nil {
			return nil, err
		}
		
		var ism types.AbstractIsm
		err = k.cdc.UnmarshalInterface(iterator.Value(), &ism)
		if err != nil {
			return nil, err
		}
		
		ismAny, err := types.PackAbstractIsm(ism)
		if err != nil {
			return nil, err
		}

		customIsms = append(customIsms, &types.CustomIsm{
			Index: uint32(index64),
			AbstractIsm: ismAny,
		})
	}

	return customIsms, nil
}

// storeCustomIsm store the custom ISM
func (k Keeper) storeCustomIsm(ctx sdk.Context, index uint32, ism types.AbstractIsm) error {
	store := ctx.KVStore(k.storeKey)

	ismBz, err := k.cdc.MarshalInterface(ism)
	if err != nil {
		return err
	}
	store.Set(types.CustomIsmKey(index), ismBz)

	return nil
}

// getNextCustomIsmIndex gets the next index to be used
func (k Keeper) getNextCustomIsmIndex(ctx sdk.Context) (uint32, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte(types.KeyCustomIsm))
	defer iterator.Close()

	index := uint32(1)
	if iterator.Valid() {
		indexBz := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyCustomIsm)))
		index64, err := strconv.ParseUint(string(indexBz), 10, 32)
		if err != nil {
			return 0, err
		}
		index64++
		index = uint32(index64)
	}

	return index, nil
}