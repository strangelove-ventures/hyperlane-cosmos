package message_id_multisig

// Message ID Multisig metadata
const (
	ORIGIN_MAILBOX_OFFSET = 0
	MERKLE_ROOT_OFFSET    = 32
	SIGNATURES_OFFSET     = 64
	SIGNATURE_LENGTH      = 65
)

func OriginMailbox(metadata []byte) []byte {
	return metadata[ORIGIN_MAILBOX_OFFSET:MERKLE_ROOT_OFFSET]
}

func Root(metadata []byte) []byte {
	return metadata[MERKLE_ROOT_OFFSET:SIGNATURES_OFFSET]
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
