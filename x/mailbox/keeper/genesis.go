package keeper

import (
	"bytes"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

// TODO: replace tracking each item with tracking only the count and branches (major improvement)

// InitGenesis initializes the hyperlane mailbox module's state from a provided genesis
// state.
func (k *Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) error {

	k.ImtCount = gs.Tree.Count

	for _, msgDelivered := range gs.DeliveredMessages {
		k.Delivered[msgDelivered.Id] = true
	}

	k.SetDomain(ctx, gs.Domain)
	return nil
}

// ExportGenesis returns the hyperlane mailbox module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	store := ctx.KVStore(k.storeKey)

	var genesisState types.GenesisState
	genesisState.Tree.Count = k.ImtCount
	genesisState.Tree.TreeEntries = ExportTreeEntries(store)
	genesisState.DeliveredMessages = ExportDeliveredGenesis(store)
	return genesisState
}

func ExportTreeEntries(store sdk.KVStore) []*types.TreeEntry {
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyMailboxIMT))
	defer iterator.Close()

	var treeEntries []*types.TreeEntry
	prefix := []byte(fmt.Sprintf("%s/", types.KeyMailboxIMT))

	for ; iterator.Valid(); iterator.Next() {
		indexBytes := bytes.TrimPrefix(iterator.Key(), prefix)
		index, err := strconv.ParseUint(string(indexBytes), 10, 32)
		if err != nil {
			panic(err)
		}
		treeEntries = append(treeEntries, &types.TreeEntry{
			Index:   uint32(index),
			Message: iterator.Value(),
		})
	}
	return treeEntries
}

func ExportDeliveredGenesis(store sdk.KVStore) []*types.MessageDelivered {
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
