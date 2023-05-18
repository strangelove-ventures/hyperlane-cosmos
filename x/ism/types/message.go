package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/crypto"
)

// This should be in a common library

const (
	VERSION_OFFSET     = 0
	NONCE_OFFSET       = 1
	ORIGIN_OFFSET      = 5
	SENDER_OFFSET      = 9
	DESTINATION_OFFSET = 41
	RECIPIENT_OFFSET   = 45
	BODY_OFFSET        = 77
)

// id
// nonce
// origin
func Id(message []byte) []byte {
	return crypto.Keccak256(message)
}

func Nonce(message []byte) uint32 {
	return binary.BigEndian.Uint32(message[NONCE_OFFSET:ORIGIN_OFFSET])
}

func Origin(message []byte) uint32 {
	return binary.BigEndian.Uint32(message[ORIGIN_OFFSET:SENDER_OFFSET])
}
