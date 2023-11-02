package receiver

type Receiver interface {
	// Address() returns the bech32 address of module
	Address() string
	// QueryIsm returns a custom ism id or 0 for default
	QueryIsm() uint32
	// Process() will process the message, sender is in hex
	Process(origin uint32, sender, msg string) error
}