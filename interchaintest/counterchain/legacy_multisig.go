package counterchain

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	imt "github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
)

// Validator will not have proof -- needs to get it from relayer which assembles it. Relayer assembles message, metadata, and proof
func (c *CounterChain) CreateLegacyMetadata(message []byte, proof [imt.TreeDepth][32]byte) (metadata []byte) {
	require.Equal(c.T, LEGACY_MULTISIG, c.IsmType)

	merkleRoot := c.Tree.Root() // comes from origin, this should be queried from the simd mailbox QueryCurrentTreeMetadataResponse
	metadata = append(metadata, merkleRoot...)

	// comes from origin, this should be queried from the simd mailbox QueryCurrentTreeMetadataResponse
	index := common.Nonce(message)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index)
	metadata = append(metadata, indexBytes...)

	originMailbox, err := hex.DecodeString("000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3")
	if err != nil {
		panic(err)
	}
	metadata = append(metadata, originMailbox...)

	for i := 0; i < imt.TreeDepth; i++ {
		metadata = append(metadata, proof[i][:]...)
	}

	metadata = append(metadata, c.ValSet.Threshold)

	// Validator functionality is below this line ...
	// Validator will need to query the origin chain to get the merkle root and index
	checkpoint := legacy_multisig.Digest(common.Origin(message), originMailbox, merkleRoot, index)
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(checkpoint, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		metadata = append(metadata, sig...)
	}

	return metadata
}

// Relayer assembles message, metadata, and proof but does not sign
func (c *CounterChain) CreateRelayerLegacyMetadata(message []byte, proof [imt.TreeDepth][32]byte, originMailbox []byte) (metadata []byte) {
	require.Equal(c.T, LEGACY_MULTISIG, c.IsmType)

	merkleRoot := c.Tree.Root() // comes from origin, this should be queried from the simd mailbox QueryCurrentTreeMetadataResponse
	metadata = append(metadata, merkleRoot...)

	// comes from origin, this should be queried from the simd mailbox QueryCurrentTreeMetadataResponse
	index := common.Nonce(message)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index)
	metadata = append(metadata, indexBytes...)
	metadata = append(metadata, originMailbox...)

	for i := 0; i < imt.TreeDepth; i++ {
		metadata = append(metadata, proof[i][:]...)
	}

	metadata = append(metadata, c.ValSet.Threshold)
	return metadata
}

func (c *CounterChain) Sign(digest []byte) []byte {
	var signature []byte
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(digest, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		signature = append(signature, sig...)
	}
	return signature
}
