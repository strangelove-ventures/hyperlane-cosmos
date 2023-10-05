package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
	mailbox "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/keeper"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer
	storeKey      storetypes.StoreKey
	cdc           codec.BinaryCodec
	mailboxKeeper mailbox.ReadOnlyMailboxKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, mailboxKeeper mailbox.ReadOnlyMailboxKeeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		mailboxKeeper: mailboxKeeper,
	}
}
