package types

type MailboxMsgType struct {
	MsgDispatch *MsgDispatch `json:"msg_dispatch,omitempty"`
}

// MsgDispatch defines the request type for the Dispatch rpc.
type MsgDispatch struct {
	Sender            string `json:"sender,omitempty"`
	DestinationDomain uint32 `json:"destination_domain,omitempty"`
	RecipientAddress  string `json:"recipient_address,omitempty"`
	MessageBody       string `json:"message_body,omitempty"`
}