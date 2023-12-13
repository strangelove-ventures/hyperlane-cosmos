package legacy_multisig

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

// RegisterInterfaces registers the LegacyMultiSig
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*types.AbstractIsm)(nil),
		&LegacyMultiSig{},
	)
}
