package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

// getGasPaidStore gets the gasPaid store for the given destination chain.
func (k Keeper) getGasPaidStore(ctx sdk.Context, destination uint32, relayer sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	destinationDomain := make([]byte, 4)
	binary.LittleEndian.PutUint32(destinationDomain, destination)
	domainKey := append(types.GasPaidKey, destinationDomain...)
	return prefix.NewStore(store, append(domainKey, relayer...))
}

// getDefaultRelayer gets the global default relayer. Note: at present, relayers are not configured per-destination.
func (k Keeper) getDefaultRelayer(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.DefaultRelayerKey)
}
