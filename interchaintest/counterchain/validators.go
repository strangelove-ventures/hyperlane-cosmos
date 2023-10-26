package counterchain

import (
	"crypto/ecdsa"
	"testing"

	//"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

type ValSet struct {
	Vals      []Val
	Total     uint32
	Threshold uint8
}

type Val struct {
	Addr string
	Priv *ecdsa.PrivateKey
}

func CreateVal(t *testing.T) *Val {
	priv, err := crypto.GenerateKey()
	require.NoError(t, err)

	signer := hexutil.Encode(crypto.PubkeyToAddress(priv.PublicKey).Bytes())

	return &Val{
		Addr: signer,
		Priv: priv,
	}
}

func CreateValFromKey(t *testing.T, ecdsaPrivKey string) *Val {
	priv, err := crypto.HexToECDSA(ecdsaPrivKey)
	require.NoError(t, err)

	signer := hexutil.Encode(crypto.PubkeyToAddress(priv.PublicKey).Bytes())

	return &Val{
		Addr: signer,
		Priv: priv,
	}
}

func CreateValSet(t *testing.T, total uint32, threshold uint8) *ValSet {
	var valSet ValSet
	for i := uint32(0); i < total; i++ {
		valSet.Vals = append(valSet.Vals, *CreateVal(t))
	}

	valSet.Threshold = threshold
	valSet.Total = total

	return &valSet
}
