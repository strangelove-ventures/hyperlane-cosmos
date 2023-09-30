package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: key,
	}
}
