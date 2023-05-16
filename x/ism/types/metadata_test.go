package types_test

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"testing"

<<<<<<< HEAD
	"github.com/stretchr/testify/require"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

type MultisigIsmMetadata struct {
	Root          []byte
	Index         uint32
	OriginMailbox []byte
	Proof         []byte
	signatures    [][]byte
=======
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/stretchr/testify/require"
)

type MultisigIsmMetadata struct {
	Root []byte
	Index uint32
	OriginMailbox []byte
	Proof []byte
	signatures [][]byte
>>>>>>> da5a6f7... Verify merkle proof working and verify validator signature WIP
}

func TestMetadataSuccess(t *testing.T) {
	var metadata []byte
	root, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000011")
	require.NoError(t, err)
	metadata = append(metadata, root...)
	index := uint32(1)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index)
	metadata = append(metadata, indexBytes...)
	originMailbox, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000012")
	require.NoError(t, err)
	metadata = append(metadata, originMailbox...)
	proof := make([]byte, 1023)
	temp, err := hex.DecodeString("13")
	proof = append(proof, temp...)
	require.NoError(t, err)
	metadata = append(metadata, proof...)
	var signatures [][]byte
	prefixSig := make([]byte, 64)
<<<<<<< HEAD
	for i := 0; i < 5; i++ {
		signature, err := hex.DecodeString(fmt.Sprintf("2%d", i))
=======
	for i := 0; i<5; i++ {
		signature, err := hex.DecodeString(fmt.Sprintf("2%d",i))
>>>>>>> da5a6f7... Verify merkle proof working and verify validator signature WIP
		require.NoError(t, err)
		signature = append(prefixSig, signature...)
		signatures = append(signatures, signature)
		metadata = append(metadata, signature...)
	}

	getRoot := types.Root(metadata)
	require.Equal(t, root, getRoot)

	getIndex := types.Index(metadata)
	require.Equal(t, index, getIndex)

	getOriginMailbox := types.OriginMailbox(metadata)
	require.Equal(t, originMailbox, getOriginMailbox)

	getProof := types.Proof(metadata)
	require.Equal(t, proof, getProof)

<<<<<<< HEAD
	for i := 0; i < 5; i++ {
		getSig := types.SignatureAt(metadata, uint32(i))
		require.Equal(t, signatures[i], getSig)
	}
}
=======
	for i := 0; i<5; i++ {
		getSig := types.SignatureAt(metadata, uint32(i))
		require.Equal(t, signatures[i], getSig)
	}	
}
>>>>>>> da5a6f7... Verify merkle proof working and verify validator signature WIP
