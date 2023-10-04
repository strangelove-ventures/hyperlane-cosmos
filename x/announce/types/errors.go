package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrPackAnnouncement    = errorsmod.Register(ModuleName, 1, "failed packing announcement")
	ErrMarshalAnnouncement = errorsmod.Register(ModuleName, 2, "failed marshalling announcement")
	ErrInvalidValidator    = errorsmod.Register(ModuleName, 3, "invalid validator address")
)
