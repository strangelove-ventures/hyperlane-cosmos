package types

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

// This should be in a common library

func Digest(origin uint32, originMailbox []byte, root []byte, index uint32) []byte {
	domainHash := DomainHash(origin, originMailbox)

	var packed []byte
	packed = append(packed, domainHash...)
	packed = append(packed, root...)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index)
	packed = append(packed, indexBytes...)

	packedHash := crypto.Keccak256(packed)
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(packedHash), packedHash)
	return crypto.Keccak256([]byte(msg))
}

func DomainHash(origin uint32, originMailbox []byte) []byte {
	var packed []byte

	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	packed = append(packed, originBytes...)
	packed = append(packed, originMailbox...)
	packed = append(packed, []byte("HYPERLANE")...)

	return crypto.Keccak256(packed)
}
