package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var ErrInvalid = sdkerrors.Register(ModuleName, 1, "invalid")
var ErrMsgTooLong = sdkerrors.Register(ModuleName, 2, "msg too long")
