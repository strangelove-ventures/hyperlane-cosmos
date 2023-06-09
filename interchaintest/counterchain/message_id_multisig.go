package counterchain

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
)

func (c *CounterChain) CreateMessageIdMetadata(message []byte) (metadata []byte) {
	originMailbox := []byte("12345678901234567890123456789012") // Shouldn't matter
	metadata = append(metadata, originMailbox...)

	merkleRoot := c.Tree.Root()
	metadata = append(metadata, merkleRoot...)

	// Isn't this supposed to be the checkpoint?
	id := common.Id(message)
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(id, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		metadata = append(metadata, sig...)
	}

	return metadata
}