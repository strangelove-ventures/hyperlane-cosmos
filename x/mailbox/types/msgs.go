package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var (
	_ sdk.Msg = (*MsgDispatch)(nil)
	_ sdk.Msg = (*MsgProcess)(nil)
)

// NewMsgDispatch creates a new MsgStoreCode instance
//
//nolint:interfacer
func NewMsgDispatch(sender string, destinationDomain uint32, recipientAddress string, messageBody string) *MsgDispatch {
	return &MsgDispatch{
		Sender:            sender,
		DestinationDomain: destinationDomain,
		RecipientAddress:  recipientAddress,
		MessageBody:       messageBody,
	}
}

// ValidateBasic implements sdk.Msg
func (m MsgDispatch) ValidateBasic() error {
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return ErrMsgDispatchInvalidSender
	}
	// Verify destination domain != 0
	if m.DestinationDomain == 0 {
		return ErrMsgDispatchInvalidDomain
	}
	// Verify recipient address is in hex with a "0x" prefix
	_, err = hexutil.Decode(m.RecipientAddress)
	if err != nil {
		return ErrMsgDispatchInvalidRecipient
	}
	// Verify message body is in hex with a "0x" prefix
	_, err = hexutil.Decode(m.MessageBody)
	if err != nil {
		return ErrMsgDispatchInvalidMsgBody
	}

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
	// Verify sender is a valid bech32 address
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return ErrMsgProcessInvalidSender
	}
	// Verify metadata
	if len(m.Metadata) < ismtypes.SIGNATURES_OFFSET {
		return ErrMsgProcessInvalidMetadata
	}
	// Verify message
	if len(m.Message) < ismtypes.BODY_OFFSET {
		return ErrMsgProcessInvalidMessage
	}
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
