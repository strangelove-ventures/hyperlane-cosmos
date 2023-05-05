package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgSetDefaultIsm)(nil)

// MsgStoreCode creates a new MsgStoreCode instance
//
//nolint:interfacer
func NewMsgSetDefaultIsm(signer string, validator_pub_keys [][]byte, threshold uint32) *MsgSetDefaultIsm {
	return &MsgSetDefaultIsm{
		Signer:  signer,
		ValidatorPubKeys: validator_pub_keys,
		Threshold: threshold,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetDefaultIsm) ValidateBasic() error {
	if len(m.ValidatorPubKeys) == 0 {
		return ErrInvalidValSet
	}
	if m.Threshold == 0 {
		return ErrInvalidThreshold
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
