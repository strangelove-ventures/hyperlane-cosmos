package types

const (
	EventTypeAnnounce         = "ValidatorAnnouncement"
	AttributeKeySender        = "sender"
	AttributeValidatorAddress = "address"  // last 20 bytes of the keccak256 hash of the validator public key
	AttributeStorageLocation  = "location" // announced storage location for the validator signatures
)
