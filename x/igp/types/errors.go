package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrGasPaid        = sdkerrors.Register(ModuleName, 2, "message gas already paid")
	ErrInvalidRelayer = sdkerrors.Register(ModuleName, 3, "invalid relayer")
)
