package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrGasPaid                 = errorsmod.Register(ModuleName, 2, "message gas already paid")
	ErrInvalidRelayer          = errorsmod.Register(ModuleName, 3, "invalid relayer")
	ErrOracleUnauthorized      = errorsmod.Register(ModuleName, 4, "unauthorized to set oracle configuration")
	ErrInvalidIgp              = errorsmod.Register(ModuleName, 5, "invalid IGP")
	ErrBeneficiaryUnauthorized = errorsmod.Register(ModuleName, 6, "unauthorized to set beneficiary")
	ErrExchangeRateScale       = errorsmod.Register(ModuleName, 7, "invalid exchange rate scale")
	ErrInvalidPaymentDenom     = errorsmod.Register(ModuleName, 8, "invalid payment denom")
	ErrEmptyMessageId          = errorsmod.Register(ModuleName, 9, "invalid message ID (empty string)")
	ErrInvalidOracleConfig     = errorsmod.Register(ModuleName, 10, "invalid oracle configuration (empty)")
	ErrExchangeRate            = errorsmod.Register(ModuleName, 11, "invalid token exchange rate")
	ErrGasPrice                = errorsmod.Register(ModuleName, 12, "invalid gas price, must be gt zero")
	ErrGasOverhead             = errorsmod.Register(ModuleName, 13, "invalid gas overhead, must be gte zero")
)
