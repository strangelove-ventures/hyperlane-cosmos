package keeper

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

// InitGenesis initializes the hyperlane mailbox module's state from a provided genesis
// state.
func (k *Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {
	k.StoreTree(ctx, gs.Tree)

	store := ctx.KVStore(k.storeKey)
	// Delivered Messages.
	for _, msgDelivered := range gs.DeliveredMessages {
		store.Set(types.MailboxDeliveredKey(msgDelivered.Id), []byte{1})
	}
	// Domain
	k.SetDomain(ctx, gs.Domain)
	return nil
}

// ExportGenesis returns the hyperlane mailbox module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	tree := k.GetTree(ctx)

	return types.GenesisState{
		DeliveredMessages: ExportDeliveredMessages(ctx.KVStore(k.storeKey)),
		Tree:              tree,
		Domain:            k.GetDomain(ctx),
	}
}

func ExportDeliveredMessages(store sdk.KVStore) []*types.MessageDelivered {
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyMailboxDelivered))
	defer iterator.Close()

	var delivered []*types.MessageDelivered
	prefix := []byte(fmt.Sprintf("%s/", types.KeyMailboxDelivered))

	for ; iterator.Valid(); iterator.Next() {
		idBytes := bytes.TrimPrefix(iterator.Key(), prefix)
		delivered = append(delivered, &types.MessageDelivered{
			Id: string(idBytes),
		})
	}
	return delivered
}
