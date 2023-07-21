package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer
	storeKey    storetypes.StoreKey
	cdc         codec.BinaryCodec
	gasoracles  map[uint32]types.GasOracleConfig
	beneficiary string
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, beneficiary string) Keeper {
	return Keeper{
		cdc:         cdc,
		storeKey:    key,
		gasoracles:  map[uint32]types.GasOracleConfig{},
		beneficiary: beneficiary,
	}
}
