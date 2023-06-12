package types

import (
	"fmt"
)

const (
	// ModuleName for the hyperlane mailbox
	ModuleName = "hyperlane-ism"

	// StoreKey is the store key string for hyperlane mailbox
	StoreKey = ModuleName

	KeyOriginsDefaultIsm = "defaultIsm"
)

func OriginKey(origin uint32) []byte {
	return []byte(fmt.Sprintf("%s/%d", KeyOriginsDefaultIsm, origin))
}
