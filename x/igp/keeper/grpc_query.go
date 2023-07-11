package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// GetBeneficiary implements the Query
func (k Keeper) GetBeneficiary(c context.Context, req *types.GetBeneficiaryRequest) (*types.GetBeneficiaryResponse, error) {
	if req == nil || *req == (types.GetBeneficiaryRequest{}) {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	return &types.GetBeneficiaryResponse{}, nil
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
