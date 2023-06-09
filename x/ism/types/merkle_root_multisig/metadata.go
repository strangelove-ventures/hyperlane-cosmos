package merkle_root_multisig

import (
	"encoding/binary"
)

// Merkle root multisig metadata released with message id multisig
// Note: Not currently used yet, see legacy folder for heckpoint metadata in use
const (
	ORIGIN_MAILBOX_OFFSET        = 0
	CHECKPOINT_INDEX_OFFSET      = 32
	CHECKPOINT_MESSAGE_ID_OFFSET = 36
	MERKLE_PROOF_OFFSET          = 68
	MERKLE_PROOF_LENGTH          = 32 * 32
	SIGNATURES_OFFSET            = 1092
	SIGNATURE_LENGTH             = 65
)

func OriginMailbox(metadata []byte) []byte {
	return metadata[ORIGIN_MAILBOX_OFFSET:ORIGIN_MAILBOX_OFFSET+32]
}

func Index(metadata []byte) uint32 {
	return binary.BigEndian.Uint32(metadata[CHECKPOINT_INDEX_OFFSET:CHECKPOINT_INDEX_OFFSET+4])
}

func MessageId(metadata []byte) []byte {
	return metadata[CHECKPOINT_MESSAGE_ID_OFFSET:CHECKPOINT_MESSAGE_ID_OFFSET+32]
}

func Proof(metadata []byte) []byte {
	return metadata[MERKLE_PROOF_OFFSET:MERKLE_PROOF_OFFSET + MERKLE_PROOF_LENGTH]
}

func SignatureAt(metadata []byte, index uint32) []byte {
	start := SIGNATURES_OFFSET + (index * SIGNATURE_LENGTH)
	end := start + SIGNATURE_LENGTH
	signature := metadata[start:end]
	if signature[64] >= 4 {
		signature[64] = signature[64] - 27
	}
	return signature
}
