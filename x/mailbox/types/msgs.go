package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgSetDefaultIsm)(nil)

// MsgStoreCode creates a new MsgStoreCode instance
//
//nolint:interfacer
func NewMsgSetDefaultIsm(signer string, ismHash string) *MsgSetDefaultIsm {
	return &MsgSetDefaultIsm{
		Signer:  signer,
		IsmHash: ismHash,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetDefaultIsm) ValidateBasic() error {
	if len(m.IsmHash) == 0 {
		return ErrInvalidIsmHash
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
