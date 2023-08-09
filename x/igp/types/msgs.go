package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgCreateIgp)(nil)
	_ sdk.Msg = (*MsgPayForGas)(nil)
	_ sdk.Msg = (*MsgClaim)(nil)
	_ sdk.Msg = (*MsgSetGasOracles)(nil)
	_ sdk.Msg = (*MsgSetBeneficiary)(nil)
	_ sdk.Msg = (*MsgSetRemoteGasData)(nil)
	_ sdk.Msg = (*MsgSetDestinationGasOverhead)(nil)
)

// GetSigners implements sdk.Msg
func (m MsgSetDestinationGasOverhead) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetDestinationGasOverhead) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	if m.GasOverhead.LT(math.ZeroInt()) {
		return ErrGasOverhead
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m MsgSetRemoteGasData) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetRemoteGasData) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	if m.TokenExchangeRate.LTE(math.ZeroInt()) {
		return ErrExchangeRate
	}

	if m.GasPrice.LTE(math.ZeroInt()) {
		return ErrGasPrice
	}
	return nil
}

// GetSigners implements sdk.Msg
func (m MsgCreateIgp) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgCreateIgp) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	// Verify address is a valid bech32 address
	_, err = sdk.AccAddressFromBech32(m.Beneficiary)
	if err != nil {
		return err
	}

	if !m.TokenExchangeRateScale.IsZero() && m.TokenExchangeRateScale.LT(math.OneInt()) {
		return ErrExchangeRateScale
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m MsgSetBeneficiary) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetBeneficiary) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	// Verify address is a valid bech32 address
	_, err = sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return err
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m MsgSetGasOracles) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgSetGasOracles) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}

	if len(m.Configs) == 0 {
		return ErrInvalidOracleConfig
	}

	for _, conf := range m.Configs {
		_, err := sdk.AccAddressFromBech32(conf.GasOracle)
		if err != nil {
			return err
		}
	}

	return nil
}

// ValidateBasic implements sdk.Msg
func (m MsgPayForGas) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}
	if m.MessageId == "" {
		return ErrEmptyMessageId
	}
	if err := m.MaximumPayment.Validate(); err != nil {
		return err
	}

	return nil
}

// GetSigners implements sdk.Msg
func (m MsgPayForGas) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}

// ValidateBasic implements sdk.Msg
func (m MsgClaim) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	return err
}

// GetSigners implements sdk.Msg
func (m MsgClaim) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{signer}
}
