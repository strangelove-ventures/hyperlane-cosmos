package keeper

import (
	"context"

	"cosmossdk.io/math"
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
func (k Keeper) QuoteGasPayment(ctx context.Context, req *types.QuoteGasPaymentRequest) (*types.QuoteGasPaymentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	igp, err := k.getIgp(sdkCtx, req.IgpId)
	if err != nil {
		return nil, err
	}

	oracle, ok := igp.Oracles[req.DestinationDomain]
	if !ok {
		return nil, types.ErrOracleUnauthorized.Wrapf("oracle with destination %d does not exist for IGP %d", req.DestinationDomain, igp.IgpId)
	}

	destGasAmount := k.getDestinationGasAmount(oracle, req.GasAmount)

	// It's possible the IGP creator did not set the scale.
	exchRateScale := igp.TokenExchangeRateScale
	if exchRateScale.IsZero() {
		exchRateScale = math.OneInt()
	}

	exchRate := oracle.TokenExchangeRate
	gasPrice := oracle.GasPrice
	destGasCost := destGasAmount.Mul(gasPrice)
	nativePrice := destGasCost.Mul(exchRate).Quo(exchRateScale)

	return &types.QuoteGasPaymentResponse{
		Amount: nativePrice,
		Denom:  k.stakingKeeper.BondDenom(sdkCtx),
	}, nil
}

// GetExchangeRateAndGasPrice implements the Query
func (k Keeper) GetExchangeRateAndGasPrice(ctx context.Context, req *types.GetExchangeRateAndGasPriceRequest) (*types.GetExchangeRateAndGasPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	//TODO: replace with actual IGP from param
	igp_ph := uint32(0)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	igp, err := k.getIgp(sdkCtx, igp_ph)
	if err != nil {
		return nil, err
	}

	oracle, ok := igp.Oracles[req.DestinationDomain]
	if !ok {
		return nil, types.ErrOracleUnauthorized.Wrapf("oracle with destination %d does not exist for IGP %d", req.DestinationDomain, igp.IgpId)
	}

	return &types.GetExchangeRateAndGasPriceResponse{TokenExchangeRate: oracle.TokenExchangeRate, GasPrice: oracle.GasPrice}, nil
}
