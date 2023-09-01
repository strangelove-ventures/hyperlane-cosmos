package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalid                   = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrInvalidOriginIsm          = sdkerrors.Register(ModuleName, 2, "no ism found for origin")
	ErrInvalidIsmSet             = sdkerrors.Register(ModuleName, 3, "invalid ism set")
	ErrInvalidValSet             = sdkerrors.Register(ModuleName, 4, "invalid val set")
	ErrInvalidThreshold          = sdkerrors.Register(ModuleName, 5, "invalid threshold")
	ErrPackAny                   = sdkerrors.Register(ModuleName, 6, "failed packing ism to any")
	ErrUnpackAny                 = sdkerrors.Register(ModuleName, 7, "failed unpacking ism from any")
	ErrVerifyMerkleProof         = sdkerrors.Register(ModuleName, 8, "verify merkle proof failed")
	ErrVerifyValidatorSignatures = sdkerrors.Register(ModuleName, 9, "verify validator signatures failed")
)
