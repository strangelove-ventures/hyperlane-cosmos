package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	proto "github.com/cosmos/gogoproto/proto"
)

// RegisterInterfaces registers the hyperlane mailbox
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"hyperlane.ism.v1.AbstractIsm",
		(*AbstractIsm)(nil),
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSetDefaultIsm{},
		&MsgCreateIsm{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func PackAbstractIsm(ism AbstractIsm) (*codectypes.Any, error) {
	msg, ok := ism.(proto.Message)
	if !ok {
		return nil, sdkerrors.Wrapf(ErrPackAny, "cannot proto marshal %T", ism)
	}

	anyClientState, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrPackAny, err.Error())
	}

	return anyClientState, nil
}

func MustPackAbstractIsm(ism AbstractIsm) *codectypes.Any {
	anyIsm, err := PackAbstractIsm(ism)
	if err != nil {
		panic(err)
	}
	return anyIsm
}

func UnpackAbstractIsm(any *codectypes.Any) (AbstractIsm, error) {
	if any == nil {
		return nil, sdkerrors.Wrap(ErrUnpackAny, "protobuf Any message cannot be nil")
	}

	ism, ok := any.GetCachedValue().(AbstractIsm)
	if !ok {
		return nil, sdkerrors.Wrapf(ErrUnpackAny, "cannot unpack Any into Ism %T", any)
	}

	return ism, nil
}

func MustUnpackAbstractIsm(any *codectypes.Any) AbstractIsm {
	ism, err := UnpackAbstractIsm(any)
	if err != nil {
		panic(err)
	}
	return ism
}
