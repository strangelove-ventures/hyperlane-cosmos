package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalid               = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrMsgTooLong            = sdkerrors.Register(ModuleName, 2, "msg too long")
	ErrMsgInvalidVersion     = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrMsgInvalidDomain      = sdkerrors.Register(ModuleName, 4, "invalid domain")
	ErrMsgDelivered          = sdkerrors.Register(ModuleName, 5, "message delivered")
	ErrMsgVerificationFailed = sdkerrors.Register(ModuleName, 6, "message verification failed")
	ErrMsgHandling           = sdkerrors.Register(ModuleName, 7, "message failed during handle")
)
