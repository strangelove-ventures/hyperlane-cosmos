package keeper

import (
	"cosmossdk.io/math"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

// getDestinationGasOverhead Returns the amount+gas overhead for the oracle
func (k Keeper) getDestinationGasAmount(oracle *types.GasOracle, gasAmount math.Int) math.Int {
	return oracle.GasOverhead.Add(gasAmount)
}
