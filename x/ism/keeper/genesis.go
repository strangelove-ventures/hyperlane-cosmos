package keeper

import (
	"bytes"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

// InitGenesis initializes the hyperlane ISM module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	for _, originIsm := range gs.DefaultIsm {
		ism, err := types.UnpackAbstractIsm(originIsm.AbstractIsm)
		if err != nil {
			return err
		}
		err = k.storeDefaultIsm(ctx, originIsm.Origin, ism)
		if err != nil {
			return err
		}
	}
	for _, customIsm := range gs.CustomIsm {
		ism, err := types.UnpackAbstractIsm(customIsm.AbstractIsm)
		if err != nil {
			return err
		}
		err = k.storeCustomIsm(ctx, customIsm.Index, ism)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExportGenesis returns the hyperlane ISM module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	return types.GenesisState{
		DefaultIsm: k.ExportDefaultIsms(ctx),
		CustomIsm:  k.ExportCustomIsms(ctx),
	}
}

// ExportDefaultIsms return the default ISMs
func (k Keeper) ExportDefaultIsms(ctx sdk.Context) []types.DefaultIsm {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyOriginsDefaultIsm))
	defer iterator.Close()

	var defaultIsms []types.DefaultIsm
	for ; iterator.Valid(); iterator.Next() {
		originBz := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyOriginsDefaultIsm)))
		origin, err := strconv.ParseUint(string(originBz), 10, 32)
		if err != nil {
			panic(err)
		}
		var ism types.AbstractIsm
		err = k.cdc.UnmarshalInterface(iterator.Value(), &ism)
		if err != nil {
			panic(err)
		}
		ismAny, err := types.PackAbstractIsm(ism)
		if err != nil {
			panic(err)
		}
		defaultIsms = append(defaultIsms, types.DefaultIsm{
			Origin:      uint32(origin),
			AbstractIsm: ismAny,
		})
	}

	return defaultIsms
}

// ExportCustomIsms return the custom ISMs
func (k Keeper) ExportCustomIsms(ctx sdk.Context) []types.CustomIsm {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyCustomIsm))
	defer iterator.Close()

	var customIsms []types.CustomIsm
	for ; iterator.Valid(); iterator.Next() {
		indexBz := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyCustomIsm)))
		index, err := strconv.ParseUint(string(indexBz), 10, 32)
		if err != nil {
			panic(err)
		}
		var ism types.AbstractIsm
		err = k.cdc.UnmarshalInterface(iterator.Value(), &ism)
		if err != nil {
			panic(err)
		}
		ismAny, err := types.PackAbstractIsm(ism)
		if err != nil {
			panic(err)
		}
		customIsms = append(customIsms, types.CustomIsm{
			Index:       uint32(index),
			AbstractIsm: ismAny,
		})
	}

	return customIsms
}
