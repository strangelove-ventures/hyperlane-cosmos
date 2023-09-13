package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	authority  string
	defaultIsm map[uint32]types.AbstractIsm
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		authority:  authority,
		defaultIsm: map[uint32]types.AbstractIsm{},
	}
}

func (k Keeper) Verify(metadata, message []byte) (bool, error) {
	msgOrigin := common.Origin(message)
	// Look up recipient contract's ISM, if 0, use default multi sig (just use default for now)
	ism := k.defaultIsm[msgOrigin]
	if ism != nil {
		return ism.Verify(metadata, message)
	}
	return false, types.ErrInvalidOriginIsm.Wrapf("no ISM configured for origin %d", msgOrigin)
}
