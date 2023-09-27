package counterchain

import (
	"encoding/binary"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
)

const MAX_MESSAGE_BODY_BYTES = 2_000

var (
	LEGACY_MULTISIG      = "legacy_multisig"
	MERKLE_ROOT_MULTISIG = "merkle_root_multisig"
	MESSAGE_ID_MULTISIG  = "message_id_multisig"
)

type CounterChain struct {
	T       *testing.T
	ValSet  ValSet
	Tree    *imt.Tree
	Domain  uint32
	IsmType string
}

func CreateCounterChain(t *testing.T, domain uint32, ismType string) *CounterChain {
	return &CounterChain{
		T:       t,
		ValSet:  *CreateValSet(t, 3, 2),
		Tree:    &imt.Tree{},
		Domain:  domain,
		IsmType: ismType,
	}
}

// He's the emperor because he's the only one
func CreateEmperorValidator(t *testing.T, domain uint32, ismType string, privKey string) *CounterChain {
	var valSet ValSet
	valSet.Vals = []Val{*CreateValFromKey(t, privKey)}
	valSet.Threshold = 1
	valSet.Total = 1

	return &CounterChain{
		T:       t,
		ValSet:  valSet,
		Tree:    &imt.Tree{},
		Domain:  domain,
		IsmType: ismType,
	}
}

func (c *CounterChain) CreateMessage(sender string, originDomain uint32, destDomain uint32, recipient string, msg string) (message []byte, proof [imt.TreeDepth][32]byte) {
	version := make([]byte, 1)
	message = append(message, version...)

	// Nonce is the tree count.
	nonce := uint32(c.Tree.Count())
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, nonce)
	message = append(message, nonceBytes...)

	// Local Domain is set on NewKeeper
	origin := uint32(originDomain) // was c.Domain
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	message = append(message, originBytes...)

	// Get the Sender address
	// Remote chain is unknown, so this must be a hex string
	senderBytes := hexutil.MustDecode(sender)
	for len(senderBytes) < (common.DESTINATION_OFFSET - common.SENDER_OFFSET) {
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
	for len(recipientBytes) < (common.BODY_OFFSET - common.RECIPIENT_OFFSET) {
		padding := make([]byte, 1)
		recipientBytes = append(padding, recipientBytes...)
	}
	message = append(message, recipientBytes...)

	// Get the Message Body, will be string based
	messageBytes := []byte(msg)
	message = append(message, messageBytes...)

	// Get the message ID
	id := common.Id(message)

	proof = c.Tree.GetProofForNextIndex() // Tree corresponds to origin chain metadata

	// Insert the message id into the tree
	err := c.Tree.Insert(id)
	require.NoError(c.T, err)

	return message, proof
}
