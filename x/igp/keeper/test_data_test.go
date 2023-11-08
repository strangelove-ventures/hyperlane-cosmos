package keeper_test

import (
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
)

// Message Ids:
// 0: 41ff551fb2b112a262f2ed4b15000c70c7c4ddc07622cf97cc8013349b87c1e3 Optimism(10) -> Arbitrum(42161)
// 1: b6085ba53813d1ed7b86d37e28e705234a029d37f8e6a62ce83059b99b90031d Optimism(10) -> Ethereum(1)
// 2: 378e32da5c00d2a4121a83e024202bbec8e013817a1ae161db6029ef0a636bee Arbitrum(42161) -> Ethereum(1)
// 3: 79b43dd3bf339945a622231803fd60b2d0800283a1e465beb5f11564bb88d29b Arbitrum(42161) -> Ethereum(1)
// 4: 6fdfba470423e067d531209edecb5ab57c86c9feecdc24f78e328721e528b9cb Ethereum(1) -> Avalanche (43114)
// 5: 469dbfd80b8cad094a2fc630871f97a10819bdb7883b6d73a47b819bbb916eaa Ethereum(1) -> Avalanche (43114)

var messages = [...]string{
	"0000000e490000000a000000000000000000000000a6f0a37dfde9c2c8f46f010989c47d9edb3a9fa80000a4b100000000000000000000000096271ca0ab9eefb3ca481749c0ca4c705fd4f52348656c6c6f21",
	"0000000e2d0000000a000000000000000000000000a6f0a37dfde9c2c8f46f010989c47d9edb3a9fa8000000010000000000000000000000009311cee522a7c122b843b66cc31c6a63e2f9264148656c6c6f21",
	"0000000e660000a4b100000000000000000000000096271ca0ab9eefb3ca481749c0ca4c705fd4f523000000010000000000000000000000009311cee522a7c122b843b66cc31c6a63e2f9264148656c6c6f21",
	"0000000e490000a4b100000000000000000000000096271ca0ab9eefb3ca481749c0ca4c705fd4f523000000010000000000000000000000009311cee522a7c122b843b66cc31c6a63e2f9264148656c6c6f21",
	"0000000462000000010000000000000000000000009311cee522a7c122b843b66cc31c6a63e2f926410000a86a0000000000000000000000002a925cd8a5d919c5c6599633090c37fe38a561b648656c6c6f21",
	"000000045a000000010000000000000000000000009311cee522a7c122b843b66cc31c6a63e2f926410000a86a0000000000000000000000002a925cd8a5d919c5c6599633090c37fe38a561b648656c6c6f21",
}

var defaultIsms = []*types.DefaultIsm{
	{
		Origin: 1, // Ethereum origin
		AbstractIsm: types.MustPackAbstractIsm(
			&legacy_multisig.LegacyMultiSig{
				Threshold: 4,
				ValidatorPubKeys: []string{
					"0x4C327CCB881A7542BE77500b2833dc84c839E7b7",
					"0x892DC66F5B2f8C438E03f6323394e34A9C24F2D6",
					"0xd4C1211F0eefb97a846c4e6D6589832e52FC03db",
					"0x84cb373148eF9112b277e68Acf676Fefa9a9a9a0",
					"0x0d860C2b28BEC3af4FD3A5997283e460fF6F2789",
					"0x600c90404D5C9DF885404d2cC5350c9B314EA3A2",
				},
			},
		),
	},
	{
		Origin: 10, // Optimism origin
		AbstractIsm: types.MustPackAbstractIsm(
			&legacy_multisig.LegacyMultiSig{
				Threshold: 4,
				ValidatorPubKeys: []string{
					"0x9f2296D5cFC6b5176aDC7716C7596898dED13D35",
					"0xFA174eB2b4921bB652BC1adA3e8B00E7E280bf3c",
					"0xAFF4718d5D637466AD07441ee3B7c4aF8E328DBd",
					"0x9c10BbE8efa03a8f49DfDb5C549258e3a8dCA097",
					"0x62144D4A52a0A0335EA5BB84392ef9912461D9Dd",
					"0xC64d1EfEaB8aE222BC889fE669f75D21b23005D9",
				},
			},
		),
	},
	{
		Origin: 42161, // Arbitrum origin
		AbstractIsm: types.MustPackAbstractIsm(
			&legacy_multisig.LegacyMultiSig{
				Threshold: 4,
				ValidatorPubKeys: []string{
					"0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7",
					"0x25c6779d4610f940bF2488732E10bcFFB9D36F81",
					"0x9856dCb10fD6e5407FA74b5Ab1d3B96cc193e9b7",
					"0xD839424e2E5aCE0A81152298Dc2b1e3bB3c7fb20",
					"0xb8085c954b75B7088Bcce69E61D12fceF797cd8d",
					"0x505Dff4e0827aA5065F5E001dB888E0569D46490",
				},
			},
		),
	},
}
