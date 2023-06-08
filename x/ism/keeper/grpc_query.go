package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// OriginsDefaultIsm implements the Query origins default ISM gRPC method
func (k Keeper) OriginsDefaultIsm(c context.Context, req *types.QueryOriginsDefaultIsmRequest) (*types.QueryOriginsDefaultIsmResponse, error) {
	if req == nil || *req == (types.QueryOriginsDefaultIsmRequest{}) {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	defaultIsm := k.defaultIsm[req.Origin]
	ismAny, err := types.PackAbstractIsm(defaultIsm)
	if err != nil {
		return nil, err
	}
	return &types.QueryOriginsDefaultIsmResponse{
		DefaultIsm: ismAny,
	}, nil
}

// AllDefaultIsms implements the Query all default ISMs gRPC method
func (k Keeper) AllDefaultIsms(c context.Context, req *types.QueryAllDefaultIsmsRequest) (*types.QueryAllDefaultIsmsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var allDefaultIsms types.QueryAllDefaultIsmsResponse
	for origin := range k.defaultIsm {
		ism := k.defaultIsm[origin]
		ismAny, err := types.PackAbstractIsm(ism)
		if err != nil {
			return nil, err
		}
		allDefaultIsms.DefaultIsms = append(allDefaultIsms.DefaultIsms, &types.Ism{
			Origin: origin,
			AbstractIsm:    ismAny,
		})
	}
	return &allDefaultIsms, nil
}
