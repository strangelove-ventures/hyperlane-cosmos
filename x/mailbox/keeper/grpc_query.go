package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// CurrentTreeMetadata implements the Query/Tree gRPC method
func (k Keeper) CurrentTreeMetadata(c context.Context, req *types.QueryCurrentTreeMetadataRequest) (*types.QueryCurrentTreeMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	panic("Implement me")
}
