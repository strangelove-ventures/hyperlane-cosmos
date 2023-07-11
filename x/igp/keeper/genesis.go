package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

// InitGenesis initializes the hyperlane IGP module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	return nil
}

// ExportGenesis returns the hyperlane IGP module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	var genesisState types.GenesisState
	return genesisState
}
