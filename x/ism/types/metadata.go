package types

import (
	"encoding/binary"
)

// This should be in a common library

const (
<<<<<<< HEAD
	MERKLE_ROOT_OFFSET    = 0
	MERKLE_INDEX_OFFSET   = 32
	ORIGIN_MAILBOX_OFFSET = 36
	MERKLE_PROOF_OFFSET   = 68
	THRESHOLD_OFFSET      = 1092
	SIGNATURES_OFFSET     = 1093
	SIGNATURE_LENGTH      = 65
	// SIGNATURES_OFFSET = 1092
	// SIGNATURE_LENGTH = 65
=======
	MERKLE_ROOT_OFFSET = 0
    MERKLE_INDEX_OFFSET = 32
    ORIGIN_MAILBOX_OFFSET = 36
	MERKLE_PROOF_OFFSET = 68
	THRESHOLD_OFFSET = 1092
    SIGNATURES_OFFSET = 1093
    SIGNATURE_LENGTH = 65
    //SIGNATURES_OFFSET = 1092
    //SIGNATURE_LENGTH = 65
>>>>>>> da5a6f7... Verify merkle proof working and verify validator signature WIP
)

func Root(metadata []byte) []byte {
	return metadata[MERKLE_ROOT_OFFSET:MERKLE_INDEX_OFFSET]
}

func Index(metadata []byte) uint32 {
	return binary.BigEndian.Uint32(metadata[MERKLE_INDEX_OFFSET:ORIGIN_MAILBOX_OFFSET])
}

func OriginMailbox(metadata []byte) []byte {
	return metadata[ORIGIN_MAILBOX_OFFSET:MERKLE_PROOF_OFFSET]
}

func Proof(metadata []byte) []byte {
	return metadata[MERKLE_PROOF_OFFSET:THRESHOLD_OFFSET]
<<<<<<< HEAD
	// return metadata[MERKLE_PROOF_OFFSET:SIGNATURES_OFFSET]
=======
	//return metadata[MERKLE_PROOF_OFFSET:SIGNATURES_OFFSET]
>>>>>>> da5a6f7... Verify merkle proof working and verify validator signature WIP
}

func Threshold(metadata []byte) uint8 {
	return metadata[THRESHOLD_OFFSET:SIGNATURES_OFFSET][0]
}

func SignatureAt(metadata []byte, index uint32) []byte {
	start := SIGNATURES_OFFSET + (index * SIGNATURE_LENGTH)
	end := start + SIGNATURE_LENGTH
	signature := metadata[start:end]
	if signature[64] >= 4 {
		signature[64] = signature[64] - 27
	}
	return signature
<<<<<<< HEAD
}
=======
}
>>>>>>> da5a6f7... Verify merkle proof working and verify validator signature WIP
