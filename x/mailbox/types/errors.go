package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrInvalid = sdkerrors.Register(ModuleName, 1, "invalid")
var ErrMsgTooLong = sdkerrors.Register(ModuleName, 2, "msg too long")
var ErrMsgInvalidVersion = sdkerrors.Register(ModuleName, 3, "invalid version")
var ErrMsgInvalidDomain = sdkerrors.Register(ModuleName, 4, "invalid domain")
var ErrMsgDelivered = sdkerrors.Register(ModuleName, 5, "message delivered")
var ErrMsgVerificationFailed = sdkerrors.Register(ModuleName, 6, "message verification failed")
var ErrMsgHandling = sdkerrors.Register(ModuleName, 7, "message failed during handle")
