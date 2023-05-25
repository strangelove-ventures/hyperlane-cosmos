package types

import fmt "fmt"

const (
	// ModuleName for the hyperlane mailbox
	ModuleName = "hyperlane-mailbox"

	// StoreKey is the store key string for hyperlane mailbox
	StoreKey = ModuleName

	KeyMailboxIMT = "imt"
)

func MailboxIMTKey(address string) []byte {
	return []byte(fmt.Sprintf("%s/%s", KeyMailboxIMT, address))
}
