package types_test

import (
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	types "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
)

func TestMessageSuccess(t *testing.T) {
	var message []byte
	version := make([]byte, 1)
	message = append(message, version...)
	nonce := uint32(1234321)
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, nonce)
	message = append(message, nonceBytes...)
	origin := uint32(1122)
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	message = append(message, originBytes...)
	sender, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000011")
	require.NoError(t, err)
	message = append(message, sender...)
	destination := uint32(3344)
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, destination)
	message = append(message, destinationBytes...)
	recipient, err := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000012")
	require.NoError(t, err)
	message = append(message, recipient...)

	getNonce := types.Nonce(message)
	require.Equal(t, nonce, getNonce)

	getOrigin := types.Origin(message)
	require.Equal(t, origin, getOrigin)
}
