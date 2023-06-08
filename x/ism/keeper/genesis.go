package keeper

import (
	"bytes"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

// InitGenesis initializes the hyperlane mailbox module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	for _, originIsm := range gs.DefaultIsm {
		ism, err := types.UnpackAbstractIsm(originIsm.AbstractIsm)
		if err != nil {
			return err
		}
		k.defaultIsm[originIsm.Origin] = ism
	}
	return nil
}

// ExportGenesis returns the hyperlane mailbox module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyOriginsDefaultIsm))
	defer iterator.Close()

	var genesisState types.GenesisState
	for ; iterator.Valid(); iterator.Next() {
		originBytes := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyOriginsDefaultIsm)))
		origin, err := strconv.ParseUint(string(originBytes), 10, 32)
		if err != nil {
			panic(err)
		}
		var ism types.AbstractIsm
		err = k.cdc.UnmarshalInterface(iterator.Value(), &ism)
		if err != nil {
			panic(err)
		}
		ismAny, err := types.PackAbstractIsm(ism)
		genesisState.DefaultIsm = append(genesisState.DefaultIsm, types.Ism{
			Origin:      uint32(origin),
			AbstractIsm: ismAny,
		})
	}

	return genesisState
}
