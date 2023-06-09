package message_id_multisig_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	types "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/message_id_multisig"
)

type MultisigIsmMetadata struct {
	Root          []byte
	Index         uint32
	OriginMailbox []byte
	Proof         []byte
	signatures    [][]byte
}

func TestMetadataSuccess(t *testing.T) {
	var metadata []byte

	originMailbox, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000012")
	require.NoError(t, err)
	metadata = append(metadata, originMailbox...)

	root, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000011")
	require.NoError(t, err)
	metadata = append(metadata, root...)

	var signatures [][]byte
	prefixSig := make([]byte, 64)
	for i := 0; i < 4; i++ {
		signature, err := hex.DecodeString(fmt.Sprintf("0%d", i))
		require.NoError(t, err)
		signature = append(prefixSig, signature...)
		signatures = append(signatures, signature)
		metadata = append(metadata, signature...)
	}

	getOriginMailbox := types.OriginMailbox(metadata)
	require.Equal(t, originMailbox, getOriginMailbox)

	getRoot := types.Root(metadata)
	require.Equal(t, root, getRoot)

	for i := 0; i < 4; i++ {
		getSig := types.SignatureAt(metadata, uint32(i))
		require.Equal(t, signatures[i], getSig)
	}
}
