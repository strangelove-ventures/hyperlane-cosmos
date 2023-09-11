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
	AttributeKeyOrigin            = "origin"
	AttributeKeySender            = "sender"
	AttributeKeyDestinationDomain = "destination"
	AttributeKeyRecipientAddress  = "recipient"
	AttributeKeyMessage           = "message"
	AttributeKeyHyperlaneMessage  = "hyperlanemessage"
	AttributeKeyID                = "id"
)
