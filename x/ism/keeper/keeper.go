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

func (k Keeper) Verify(goCtx context.Context, metadata, message []byte, ismId uint32) (bool, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	msgOrigin := common.Origin(message)
	
	ism, err := k.getIsm(ctx, ismId, msgOrigin)
	if err != nil {
		return false, err
	}
	if ism != nil {
		return ism.Verify(metadata, message)
	}
	return false, types.ErrInvalidOriginIsm.Wrapf("no ISM configured for origin %d", msgOrigin)
}
