package keeper

import (
	"context"
	"encoding/binary"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// CurrentTreeMetadata implements the Query/Tree gRPC method
func (k Keeper) CurrentTreeMetadata(c context.Context, req *types.QueryCurrentTreeMetadataRequest) (*types.QueryCurrentTreeMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	tree := k.GetImtTree(c)

	return &types.QueryCurrentTreeMetadataResponse{
		Root:  tree.Root(),
		Count: tree.Count(),
	}, nil
}

func (k Keeper) CurrentTree(c context.Context, req *types.QueryCurrentTreeRequest) (*types.QueryCurrentTreeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	tree := k.GetImtTree(c)

	return &types.QueryCurrentTreeResponse{
		Branches: tree.Branch[:],
		Count:    tree.Count(),
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

// MsgDelivered implements the Query/MsgDelivered gRPC method
func (k Keeper) MsgDelivered(c context.Context, req *types.QueryMsgDeliveredRequest) (*types.QueryMsgDeliveredResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	delivered := false
	msgId := hexutil.Encode(req.MessageId)

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.MailboxDeliveredKey(msgId)) {
		delivered = true
	}

	return &types.QueryMsgDeliveredResponse{
		Delivered: delivered,
	}, nil
}

// RecipientsIsmId implements to Query/RecipientsIsmId gRPC method
func (k Keeper) RecipientsIsmId(c context.Context, req *types.QueryRecipientsIsmIdRequest) (*types.QueryRecipientsIsmIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	
	ctx := sdk.UnwrapSDKContext(c)

	recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), req.Recipient)
	ismId, err := k.getReceiversIsm(ctx, recipient)
	if err != nil {
		return nil, err
	}

	return &types.QueryRecipientsIsmIdResponse{
		IsmId: ismId,
	}, nil
}