package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalid        = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrInvalidValSet = sdkerrors.Register(ModuleName, 2, "invalid validator set")
	ErrInvalidThreshold = sdkerrors.Register(ModuleName, 3, "invalid threshold")
)
