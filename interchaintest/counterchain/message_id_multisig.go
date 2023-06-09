package counterchain

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/message_id_multisig"
)

func (c *CounterChain) CreateMessageIdMetadata(message []byte) (metadata []byte) {
	require.Equal(c.T, MESSAGE_ID_MULTISIG, c.IsmType)

	originMailbox := []byte("12345678901234567890123456789012") // Shouldn't matter
	metadata = append(metadata, originMailbox...)

	merkleRoot := c.Tree.Root()
	metadata = append(metadata, merkleRoot...)

	checkpoint := message_id_multisig.Digest(common.Origin(message), originMailbox, merkleRoot, common.Nonce(message), common.Id(message))
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(checkpoint, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		metadata = append(metadata, sig...)
	}

	return metadata
}
