package helpers

import (
	"encoding/binary"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	imt "github.com/strangelove-ventures/hyperlane-cosmos/imt"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/stretchr/testify/require"
)

const MAX_MESSAGE_BODY_BYTES = 2_000

type CounterChain struct {
	T *testing.T
	ValSet ValSet
	Tree *imt.Tree
	Domain uint32
}

func CreateCounterChain(t *testing.T, domain uint32) *CounterChain {
	return &CounterChain{
		T: t,
		ValSet: *CreateValSet(t, 3, 2),
		Tree: &imt.Tree{},
		Domain: domain,
	}
}

func (c *CounterChain) CreateMessage(sender string, destDomain uint32, recipient string, msg string) (message []byte, proof [imt.TreeDepth][32]byte) {
	version := make([]byte, 1)
	message = append(message, version...)

	// Nonce is the tree count.
	nonce := uint32(c.Tree.Count())
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, nonce)
	message = append(message, nonceBytes...)

	// Local Domain is set on NewKeeper
	origin := uint32(c.Domain)
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	message = append(message, originBytes...)

	// Get the Sender address
	// Remote chain is unknown, so this must be a hex string
	senderBytes := hexutil.MustDecode(sender)
	for len(senderBytes) < (ismtypes.DESTINATION_OFFSET - ismtypes.SENDER_OFFSET) {
		padding := make([]byte, 1)
		senderBytes = append(padding, senderBytes...)
	}
	message = append(message, senderBytes...)

	// Get the Destination Domain
	destination := destDomain
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, destination)
	message = append(message, destinationBytes...)

	// Get the Recipient address
	// Recipient is a cosmos contract address, so it must be a bech32 address
	recipientBytes := sdk.MustAccAddressFromBech32(recipient).Bytes()
	for len(recipientBytes) < (ismtypes.BODY_OFFSET - ismtypes.RECIPIENT_OFFSET) {
		padding := make([]byte, 1)
		recipientBytes = append(padding, recipientBytes...)
	}
	message = append(message, recipientBytes...)

	// Get the Message Body, will be string based
	messageBytes := []byte(msg)
	message = append(message, messageBytes...)

	// Get the message ID
	id := ismtypes.Id(message)

	// Get proof/path before insertion
	if c.Tree.Count() != uint32(0) {
		for i := 0; i<imt.TreeDepth; i++ {
			copy(proof[i][:], c.Tree.Branch[i][:])
		}
	} else {
		zeroHashes := imt.ZeroHashes()
		for i := 0; i<imt.TreeDepth; i++ {
			copy(proof[i][:], zeroHashes[i][:])
		}
	}

	// Insert the message id into the tree
	err := c.Tree.Insert(id)
	require.NoError(c.T, err)

	return message, proof
}

func (c *CounterChain) CreateMetadata(message []byte, proof [imt.TreeDepth][32]byte) (metadata []byte) {
	merkleRoot := c.Tree.Root()
	metadata = append(metadata, merkleRoot...)

	index := ismtypes.Nonce(message)
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

	id := ismtypes.Id(message)
	for i := uint8(0); i < c.ValSet.Threshold; i++ {
		sig, err := crypto.Sign(id, c.ValSet.Vals[i].Priv)
		require.NoError(c.T, err)
		metadata = append(metadata, sig...)
	}

	return metadata
}