package counterchain

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	imt "github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	legacy "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig/legacy"
)

func (c *CounterChain) CreateLegacyMetadata(message []byte, proof [imt.TreeDepth][32]byte) (metadata []byte) {
	merkleRoot := c.Tree.Root()
	metadata = append(metadata, merkleRoot...)

	index := common.Nonce(message)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index)
	metadata = append(metadata, indexBytes...)

	originMailbox := []byte("12345678901234567890123456789012") // Shouldn't matter
	metadata = append(metadata, originMailbox...)

	// How to get merkle proof?
	for i := 0; i < imt.TreeDepth; i++ {
		metadata = append(metadata, proof[i][:]...)
	}

	metadata = append(metadata, c.ValSet.Threshold)

	checkpoint := legacy.Digest(common.Origin(message), originMailbox, merkleRoot, index)
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(checkpoint, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		metadata = append(metadata, sig...)
	}

	return metadata
}