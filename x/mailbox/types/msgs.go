package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgDispatch)(nil)
	_ sdk.Msg = (*MsgProcess)(nil)
	_ sdk.Msg = (*MsgSetDefaultIsm)(nil)
)

func NewMsgDispatch(sender sdk.AccAddress, destinationDomain uint32, recipientAddress string, messageBody string) *MsgDispatch {
	return &MsgDispatch{
		Sender:            sender.String(),
		DestinationDomain: destinationDomain,
		RecipientAddress:  recipientAddress,
		MessageBody:       messageBody,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgDispatch) ValidateBasic() error {
	// TODO
	return nil
}

// GetSigners implements sdk.Msg
func (m MsgDispatch) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func NewMsgProcess(sender sdk.AccAddress, metadata string, message string) *MsgProcess {
	return &MsgProcess{
		Sender:   sender.String(),
		Metadata: metadata,
		Message:  message,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgProcess) ValidateBasic() error {
	// TODO
	return nil
}

// GetSigners implements sdk.Msg
func (m MsgProcess) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

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
