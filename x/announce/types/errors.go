package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrPackAnnouncement           = errorsmod.Register(ModuleName, 1, "failed packing announcement")
	ErrMarshalAnnouncement        = errorsmod.Register(ModuleName, 2, "failed marshalling announcement")
	ErrInvalidValidator           = errorsmod.Register(ModuleName, 3, "invalid validator address")
	ErrReplayAnnouncement         = errorsmod.Register(ModuleName, 4, "replay - validator already made this announcement")
	ErrBadDigest                  = errorsmod.Register(ModuleName, 5, "Signature does not match the declared validator")
	ErrMarshalAnnouncedValidators = errorsmod.Register(ModuleName, 6, "failed marshalling validators")
)
