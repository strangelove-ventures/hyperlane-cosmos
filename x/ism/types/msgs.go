package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgSetDefaultIsm)(nil)

// NewMsgSetDefaultIsm creates a new MsgSetDefaultIsm instance
func NewMsgSetDefaultIsm(signer string, isms []*OriginsMultiSigIsm) *MsgSetDefaultIsm {
	return &MsgSetDefaultIsm{
		Signer:  signer,
		Isms:    isms,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetDefaultIsm) ValidateBasic() error {
	if len(m.Isms) == 0 {
		return ErrInvalidIsmSet
	}
	for _, originIsm := range m.Isms {
		if originIsm.Ism.Threshold == 0 {
			return ErrInvalidThreshold
		}
		for _, validator := range originIsm.Ism.ValidatorPubKeys {
			if len(validator) != 32 {
				return ErrInvalidValSet
			}
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
