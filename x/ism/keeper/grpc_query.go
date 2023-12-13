package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// OriginsDefaultIsm implements the Query origins default ISM gRPC method
func (k Keeper) OriginsDefaultIsm(c context.Context, req *types.QueryOriginsDefaultIsmRequest) (*types.QueryOriginsDefaultIsmResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil || *req == (types.QueryOriginsDefaultIsmRequest{}) {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	defaultIsm, err := k.getDefaultIsm(ctx, req.Origin)
	if err != nil {
		return nil, err
	}
	if defaultIsm != nil {
		ismAny, err := types.PackAbstractIsm(defaultIsm)
		if err != nil {
			return nil, err
		}
		return &types.QueryOriginsDefaultIsmResponse{
			DefaultIsm: ismAny,
		}, nil
	}
	return &types.QueryOriginsDefaultIsmResponse{}, nil
}

// AllDefaultIsms implements the Query all default ISMs gRPC method
func (k Keeper) AllDefaultIsms(c context.Context, req *types.QueryAllDefaultIsmsRequest) (*types.QueryAllDefaultIsmsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	defaultIsms, err := k.getAllDefaultIsms(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryAllDefaultIsmsResponse{
		DefaultIsms: defaultIsms,
	}, nil
}

// CustomIsm implements the Query custom ISM gRPC method
func (k Keeper) CustomIsm(c context.Context, req *types.QueryCustomIsmRequest) (*types.QueryCustomIsmResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil || *req == (types.QueryCustomIsmRequest{}) {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	customIsm, err := k.getCustomIsm(ctx, req.IsmId)
	if err != nil {
		return nil, err
	}
	if customIsm != nil {
		ismAny, err := types.PackAbstractIsm(customIsm)
		if err != nil {
			return nil, err
		}
		return &types.QueryCustomIsmResponse{
			CustomIsm: ismAny,
		}, nil
	}
	return &types.QueryCustomIsmResponse{}, nil
}

// AllCustomIsms implements the Query all custom ISMs gRPC method
func (k Keeper) AllCustomIsms(c context.Context, req *types.QueryAllCustomIsmsRequest) (*types.QueryAllCustomIsmsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	customIsms, err := k.getAllCustomIsms(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryAllCustomIsmsResponse{
		CustomIsms: customIsms,
	}, nil
}
