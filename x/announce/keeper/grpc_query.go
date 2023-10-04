package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

// GetAnnouncedValidators list of validators that have been announced
func (k Keeper) GetAnnouncedValidators(ctx context.Context, req *types.GetAnnouncedValidatorsRequest) (*types.GetAnnouncedValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	_ = sdk.UnwrapSDKContext(ctx)

	return &types.GetAnnouncedValidatorsResponse{}, nil
}

// GetAnnouncedStorageLocations returns the list of storage locations for each requested validator
func (k Keeper) GetAnnouncedStorageLocations(ctx context.Context, req *types.GetAnnouncedStorageLocationsRequest) (*types.GetAnnouncedStorageLocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	_ = sdk.UnwrapSDKContext(ctx)

	return &types.GetAnnouncedStorageLocationsResponse{}, nil
}
