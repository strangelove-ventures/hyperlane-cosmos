package types

// Hyperlane Mailbox event types
const (
	EventTypeDispatch   = "dispatch"
	EventTypeDispatchId = "dispatch_id"
	EventTypeProcess    = "process"
	EventTypeProcessId  = "process_id"
)

// Hyperlane Mailbox attribute keys
const (
	AttributeKeyDestination       = "destination"
	AttributeKeyDestinationDomain = "destination"
	AttributeKeyHyperlaneMessage  = "hyperlanemessage"
	AttributeKeyID                = "id"
	AttributeKeyMessage           = "message"
	AttributeKeyNonce             = "nonce"
	AttributeKeyOrigin            = "origin"
	AttributeKeyRecipientAddress  = "recipient"
	AttributeKeySender            = "sender"
	AttributeKeyVersion           = "version"
)
