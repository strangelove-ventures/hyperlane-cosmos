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
	//Branches

	//k.Tree.Branch = append(k.Tree.Branch, gs.Tree.Branch...)
	//Delivered Messages.
	for _, msgDelivered := range gs.DeliveredMessages {
		k.Delivered[msgDelivered.Id] = true
	}
	//Domain
	k.SetDomain(ctx, gs.Domain)
	return nil
}

// ExportGenesis returns the hyperlane mailbox module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {

	return types.GenesisState{
		DeliveredMessages: ExportDeliveredMessages(ctx.KVStore(k.storeKey)),
		Tree: types.Tree{
			//Branch: k.Tree.Branch,
			Count: k.Tree.Count(),
		},
		Domain: k.GetDomain(ctx),
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
