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
	ErrEmptyMessageId          = sdkerrors.Register(ModuleName, 9, "invalid message ID (empty string)")
	ErrInvalidOracleConfig     = sdkerrors.Register(ModuleName, 10, "invalid oracle configuration (empty)")
	ErrExchangeRate            = sdkerrors.Register(ModuleName, 11, "invalid token exchange rate")
	ErrGasPrice                = sdkerrors.Register(ModuleName, 12, "invalid gas price, must be gt zero")
	ErrGasOverhead             = sdkerrors.Register(ModuleName, 13, "invalid gas overhead, must be gte zero")
)
