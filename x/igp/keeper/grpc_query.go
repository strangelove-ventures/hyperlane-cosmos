package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// GetBeneficiary returns the beneficiary for the given IGP
func (k Keeper) GetBeneficiary(ctx context.Context, req *types.GetBeneficiaryRequest) (*types.GetBeneficiaryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	igp, err := k.getIgp(sdkCtx, req.IgpId)
	if err != nil {
		return nil, err
	}

	return &types.GetBeneficiaryResponse{Address: igp.Beneficiary}, nil
}

// QuoteGasPayment implements the Query
func (k Keeper) QuoteGasPayment(c context.Context, req *types.QuoteGasPaymentRequest) (*types.QuoteGasPaymentResponse, error) {
	if req == nil || *req == (types.QuoteGasPaymentRequest{}) {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	return &types.QuoteGasPaymentResponse{}, nil
}

// GetExchangeRateAndGasPrice implements the Query
func (k Keeper) GetExchangeRateAndGasPrice(c context.Context, req *types.GetExchangeRateAndGasPriceRequest) (*types.GetExchangeRateAndGasPriceResponse, error) {
	if req == nil || *req == (types.GetExchangeRateAndGasPriceRequest{}) {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	return &types.GetExchangeRateAndGasPriceResponse{}, nil
}

// TODO: do these queries need UnpackInterfaces? See: https://github.com/cosmos/cosmos-sdk/issues/8327
