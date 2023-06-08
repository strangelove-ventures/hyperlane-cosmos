package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalid          = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrInvalidIsmSet    = sdkerrors.Register(ModuleName, 2, "invalid ism set")
	ErrInvalidValSet    = sdkerrors.Register(ModuleName, 3, "invalid val set")
	ErrInvalidThreshold = sdkerrors.Register(ModuleName, 4, "invalid threshold")
	ErrPackAny          = sdkerrors.Register(ModuleName, 5, "failed packing ism to any")
	ErrUnpackAny        = sdkerrors.Register(ModuleName, 6, "failed unpacking ism from any")
)
