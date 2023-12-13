package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalid                     = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrMsgTooLong                  = sdkerrors.Register(ModuleName, 2, "msg too long")
	ErrMsgInvalidVersion           = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrMsgInvalidDomain            = sdkerrors.Register(ModuleName, 4, "invalid domain")
	ErrMsgDelivered                = sdkerrors.Register(ModuleName, 5, "message delivered")
	ErrMsgVerificationFailed       = sdkerrors.Register(ModuleName, 6, "message verification failed")
	ErrMsgHandling                 = sdkerrors.Register(ModuleName, 7, "message failed during handle")
	ErrMsgDispatchInvalidSender    = sdkerrors.Register(ModuleName, 8, "invalid sender in msg dispatch")
	ErrMsgDispatchInvalidDomain    = sdkerrors.Register(ModuleName, 9, "invalid domain in msg dispatch")
	ErrMsgDispatchInvalidRecipient = sdkerrors.Register(ModuleName, 10, "invalid recipient in msg dispatch")
	ErrMsgDispatchInvalidMsgBody   = sdkerrors.Register(ModuleName, 11, "invalid msg body in msg dispatch")
	ErrMsgProcessInvalidSender     = sdkerrors.Register(ModuleName, 12, "invalid sender in msg process")
	ErrMsgProcessInvalidMetadata   = sdkerrors.Register(ModuleName, 13, "invalid metadata in msg process")
	ErrMsgProcessInvalidMessage    = sdkerrors.Register(ModuleName, 14, "invalid message in msg process")
	ErrInvalidRecipient            = sdkerrors.Register(ModuleName, 15, "invalid recipient")
)
