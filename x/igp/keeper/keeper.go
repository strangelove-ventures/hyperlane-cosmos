package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer
	storeKey    storetypes.StoreKey
	sendKeeper  bankTypes.SendKeeper
	cdc         codec.BinaryCodec
	gasoracles  map[uint32]types.GasOracleConfig
	beneficiary string
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, sendKeeper bankTypes.SendKeeper, beneficiary string) Keeper {
	return Keeper{
		sendKeeper:  sendKeeper,
		cdc:         cdc,
		storeKey:    key,
		gasoracles:  map[uint32]types.GasOracleConfig{},
		beneficiary: beneficiary,
	}
}
