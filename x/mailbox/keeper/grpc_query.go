package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// CurrentTreeMetadata implements the Query/Tree gRPC method
func (k Keeper) CurrentTreeMetadata(c context.Context, req *types.QueryCurrentTreeMetadataRequest) (*types.QueryCurrentTreeMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	return &types.QueryCurrentTreeMetadataResponse{
		Root:  k.Tree.Root(),
		Count: k.Tree.Count(),
	}, nil
}
