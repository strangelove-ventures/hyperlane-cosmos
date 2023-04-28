package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

// InitGenesis initializes the hyperlane mailbox module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	panic("Implement me")
}

// ExportGenesis returns the hyperlane mailbox module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	panic("Implement me")
}
