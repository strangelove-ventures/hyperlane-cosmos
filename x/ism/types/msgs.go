package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg                            = (*MsgSetDefaultIsm)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgSetDefaultIsm)(nil)
)

// NewMsgSetDefaultIsm creates a new MsgSetDefaultIsm instance
func NewMsgSetDefaultIsm(signer string, isms []*DefaultIsm) *MsgSetDefaultIsm {
	return &MsgSetDefaultIsm{
		Signer: signer,
		Isms:   isms,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetDefaultIsm) ValidateBasic() error {
	if len(m.Isms) == 0 {
		return ErrInvalidIsmSet
	}

	for _, originIsm := range m.Isms {
		ism, err := UnpackAbstractIsm(originIsm.AbstractIsm)
		if err != nil {
			return err
		}
		err = ism.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m MsgSetDefaultIsm) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (m MsgSetDefaultIsm) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var ism AbstractIsm
	for _, originIsm := range m.Isms {
		err := unpacker.UnpackAny(originIsm.AbstractIsm, &ism)
		if err != nil {
			return err
		}
	}
	return nil
}


var (
	_ sdk.Msg                            = (*MsgCreateIsm)(nil)
	_ codectypes.UnpackInterfacesMessage = (*MsgCreateIsm)(nil)
)

// NewMsgCreateIsm creates a new MsgCreateIsm instance
func NewMsgCreateIsm(signer string, ism *codectypes.Any) *MsgCreateIsm {
	return &MsgCreateIsm{
		Signer: signer,
		Ism:   ism,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgCreateIsm) ValidateBasic() error {
	if m.Ism == nil {
		return ErrInvalidIsmSet
	}

	ism, err := UnpackAbstractIsm(m.Ism)
	if err != nil {
		return err
	}
	err = ism.Validate()
	if err != nil {
		return err
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m MsgCreateIsm) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (m MsgCreateIsm) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var ism AbstractIsm
	
	err := unpacker.UnpackAny(m.Ism, &ism)
	if err != nil {
		return err
	}
	return nil
}

