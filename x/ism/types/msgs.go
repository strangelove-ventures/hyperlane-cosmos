package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var _ sdk.Msg = (*MsgSetDefaultIsm)(nil)
var _ codectypes.UnpackInterfacesMessage = (*MsgSetDefaultIsm)(nil)

// NewMsgSetDefaultIsm creates a new MsgSetDefaultIsm instance
func NewMsgSetDefaultIsm(signer string, isms []*Ism) *MsgSetDefaultIsm {
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