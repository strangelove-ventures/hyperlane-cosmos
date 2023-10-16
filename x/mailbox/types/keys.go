package types

import fmt "fmt"

const (
	// ModuleName for the hyperlane mailbox
	ModuleName = "hyperlane-mailbox"

	// StoreKey is the store key string for hyperlane mailbox
	StoreKey            = ModuleName
	KeyMailboxIMT       = "imt"
	KeyMailboxDelivered = "delivered"
)

var DomainKey = []byte{0x1}

func MailboxIMTKey() []byte {
	return []byte(KeyMailboxIMT)
}

func MailboxDeliveredKey(id string) []byte {
	return []byte(fmt.Sprintf("%s/%s", KeyMailboxDelivered, id))
}
