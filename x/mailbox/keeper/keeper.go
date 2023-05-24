package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	ism "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeKey  storetypes.StoreKey
	cdc       codec.BinaryCodec
	authority string
	version   byte
	domain    uint32
	tree      imt.Tree
	delivered map[string]bool
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, domain uint32) Keeper {
	// governance authority
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	return Keeper{
		cdc:       cdc,
		storeKey:  key,
		authority: authority.String(),
		version:   0,
		domain:    domain,
	}
}

func (k Keeper) getRecipientISM() ism.Keeper {
	panic("Implement Me")
}

func (k Keeper) HandleMessage(origin uint32, sender, recipient, body string) error {
	panic("Implement Me")
}
