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
	tempTree := make(map[uint32][]byte, gs.Tree.Count)
	for _, treeEntry := range gs.Tree.TreeEntries {
		tempTree[treeEntry.Index] = treeEntry.Message
	}
	var index uint32
	for index = 0; index < gs.Tree.Count; index++ {
		k.Tree.Insert(tempTree[index])
	}
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
	genesisState.Tree = ExportTreeGenesis(store)
	genesisState.DeliveredMessages = ExportDeliveredGenesis(store)
	return genesisState
}

func ExportTreeGenesis(store sdk.KVStore) types.Tree {
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyMailboxIMT))
	defer iterator.Close()

	var genesisTree types.Tree
	count := uint32(0)
	for ; iterator.Valid(); iterator.Next() {
		indexBytes := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyMailboxIMT)))
		index, err := strconv.ParseUint(string(indexBytes), 10, 32)
		if err != nil {
			panic(err)
		}
		genesisTree.TreeEntries = append(genesisTree.TreeEntries, &types.TreeEntry{
			Index:   uint32(index),
			Message: iterator.Value(),
		})
		count++
	}
	genesisTree.Count = count
	return genesisTree
}

func ExportDeliveredGenesis(store sdk.KVStore) []*types.MessageDelivered {
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyMailboxDelivered))
	defer iterator.Close()

	var delivered []*types.MessageDelivered
	for ; iterator.Valid(); iterator.Next() {
		idBytes := bytes.TrimPrefix(iterator.Key(), []byte(fmt.Sprintf("%s/", types.KeyMailboxDelivered)))
		delivered = append(delivered, &types.MessageDelivered{
			Id: string(idBytes),
		})
	}
	return delivered
}
