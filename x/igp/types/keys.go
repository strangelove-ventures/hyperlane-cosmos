package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName for the hyperlane interchain-gas-paymaster
	ModuleName = "hyperlane-igp"

	// StoreKey is the store key string for hyperlane interchain-gas-paymaster
	StoreKey = ModuleName
)

// KVStore keys
var (
	GasOverhead = []byte{0x00}
	GasPaidKey  = []byte{0x1}
)

func GasOverheadKey(destination uint32) []byte {
	return []byte(fmt.Sprintf("%d/%d", GasOverhead, destination))
}

func PayForGasKey(relayer sdk.AccAddress, messageId []byte) []byte {
	b := append(GasPaidKey, address.MustLengthPrefix(relayer)...)
	return append(b, messageId...)
}
