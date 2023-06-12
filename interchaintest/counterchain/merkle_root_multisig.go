package counterchain

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	imt "github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig"
)

func (c *CounterChain) CreateMerkleRootMetadata(message []byte, proof [imt.TreeDepth][32]byte) (metadata []byte) {
	require.Equal(c.T, MERKLE_ROOT_MULTISIG, c.IsmType)

	originMailbox := []byte("12345678901234567890123456789012") // Shouldn't matter
	metadata = append(metadata, originMailbox...)

	index := common.Nonce(message)
	indexBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(indexBytes, index)
	metadata = append(metadata, indexBytes...)

	merkleRoot := c.Tree.Root()
	metadata = append(metadata, merkleRoot...)

	for i := 0; i < imt.TreeDepth; i++ {
		metadata = append(metadata, proof[i][:]...)
	}

	checkpoint := merkle_root_multisig.Digest(common.Origin(message), originMailbox, merkleRoot, index, common.Id(message))
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(checkpoint, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		metadata = append(metadata, sig...)
	}

	return metadata
}
