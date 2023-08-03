package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrGasPaid                 = sdkerrors.Register(ModuleName, 2, "message gas already paid")
	ErrInvalidRelayer          = sdkerrors.Register(ModuleName, 3, "invalid relayer")
	ErrOracleUnauthorized      = sdkerrors.Register(ModuleName, 4, "unauthorized to set oracle configuration")
	ErrInvalidIgp              = sdkerrors.Register(ModuleName, 5, "invalid IGP")
	ErrBeneficiaryUnauthorized = sdkerrors.Register(ModuleName, 6, "unauthorized to set beneficiary")
	ErrExchangeRateScale       = sdkerrors.Register(ModuleName, 7, "invalid exchange rate scale")
	ErrInvalidPaymentDenom     = sdkerrors.Register(ModuleName, 8, "invalid payment denom")
)
