package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgDispatch)(nil)
	_ sdk.Msg = (*MsgProcess)(nil)
)

// NewMsgDispatch creates a new MsgStoreCode instance
//
//nolint:interfacer
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

// NewMsgProcess creates a new NewMsgProcess instance
//
//nolint:interfacer
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
