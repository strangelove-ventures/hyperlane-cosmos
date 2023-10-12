package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// GetAnnouncedValidators list of validators that have been announced
func (k Keeper) GetAnnouncedValidators(ctx context.Context, req *types.GetAnnouncedValidatorsRequest) (*types.GetAnnouncedValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	resp, err := k.getAnnouncedValidators(sdkCtx)
	if err != nil {
		return nil, types.ErrMarshalAnnouncedValidators
	}

	return resp, nil
}

// GetAnnouncedStorageLocations returns the list of storage locations for each requested validator
func (k Keeper) GetAnnouncedStorageLocations(ctx context.Context, req *types.GetAnnouncedStorageLocationsRequest) (*types.GetAnnouncedStorageLocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	resp := &types.GetAnnouncedStorageLocationsResponse{
		Metadata: []*types.StorageMetadata{},
	}

	for _, val := range req.Validator {
		md := &types.StorageMetadata{}
		announcementResp, err := k.getAnnouncements(sdkCtx, val)
		if err != nil {
			return nil, err
		}
		for _, loc := range announcementResp.Announcement {
			md.Metadata = append(md.Metadata, loc.StorageLocation)
		}
		resp.Metadata = append(resp.Metadata, md)
	}

	return resp, nil
}
