package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgAnnouncement)(nil)

// GetSigners implements sdk.Msg
func (m MsgAnnouncement) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgAnnouncement) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	if len(m.Validator) != ETHEREUM_ADDR_LEN {
		return ErrInvalidValidator.Wrapf("Validator address is %d bytes, expected %d bytes", len(m.Validator), ETHEREUM_ADDR_LEN)
	}

	return nil
}
