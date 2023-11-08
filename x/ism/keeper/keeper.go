package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, authority string) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   key,
		authority:  authority,
	}
}

func (k Keeper) Verify(goCtx context.Context, metadata, message []byte) (bool, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	msgOrigin := common.Origin(message)
	// Look up recipient contract's ISM, if 0, use default multi sig (just use default for now)
	ism, err := k.getDefaultIsm(ctx, msgOrigin)
	if err != nil {
		return false, err
	}
	if ism != nil {
		return ism.Verify(metadata, message)
	}
	return false, types.ErrInvalidOriginIsm.Wrapf("no ISM configured for origin %d", msgOrigin)
}
