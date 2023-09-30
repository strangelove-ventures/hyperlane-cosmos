package types

const (
	// ModuleName for the hyperlane interchain-gas-paymaster
	ModuleName = "hyperlane-igp"

	// StoreKey is the store key string for hyperlane interchain-gas-paymaster
	StoreKey = ModuleName
)

// KVStore keys
var (
	GasOverhead       = []byte{0x00}
	GasPaidKey        = []byte{0x1}
	DefaultRelayerKey = []byte{0x2}
	OracleKey         = []byte{0x3}
	IgpKey            = []byte{0x4}
	ClaimsKey         = []byte{0x5}
)
