package keeper

import (
	"context"
	"encoding/binary"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// Domain implements the Query/Domain gRPC method
func (k Keeper) Domain(c context.Context, req *types.QueryDomainRequest) (*types.QueryDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.DomainKey)
	domain := binary.LittleEndian.Uint32(b)

	return &types.QueryDomainResponse{
		Domain: domain,
	}, nil
}
