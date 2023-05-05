package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// DefaultIsm implements the Query default ISM gRPC method
func (k Keeper) DefaultIsm(c context.Context, req *types.QueryDefaultIsmRequest) (*types.QueryDefaultIsmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	panic("Implement me")
}

// DefaultIsm implements the Query default ISM gRPC method
func (k Keeper) ContractIsm(c context.Context, req *types.QueryContractIsmRequest) (*types.QueryContractIsmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	panic("Implement me")
}
