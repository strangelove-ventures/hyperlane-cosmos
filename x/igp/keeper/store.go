package keeper

import (
	"encoding/binary"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

// getGasPaidStore gets the gasPaid store for the given destination chain.
func (k Keeper) getGasPaidStore(ctx sdk.Context, destination uint32, relayer sdk.AccAddress) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	destinationDomain := make([]byte, 4)
	binary.LittleEndian.PutUint32(destinationDomain, destination)
	domainKey := append(types.GasPaidKey, destinationDomain...)
	return prefix.NewStore(store, append(domainKey, relayer...))
}

// getDefaultRelayer gets the global default relayer. Note: at present, relayers are not configured per-destination.
func (k Keeper) getDefaultRelayer(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.DefaultRelayerKey)
}

func (k Keeper) getIgpStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.IgpKey)
}

// getIgp unmarshal igp from storage
func (k Keeper) getIgp(ctx sdk.Context, igp_id uint32) (*types.Igp, error) {
	store := k.getIgpStore(ctx)
	igp := types.Igp{}

	igpIdB := make([]byte, 4)
	binary.LittleEndian.PutUint32(igpIdB, igp_id)
	igpB := store.Get(igpIdB)
	if igpB == nil {
		return nil, errors.New("IGP does not exist")
	}

	err := igp.Unmarshal(igpB)
	return &igp, err
}

// setIgp store the IGP
func (k Keeper) setIgp(ctx sdk.Context, igp *types.Igp) error {
	store := k.getIgpStore(ctx)
	igpIdB := make([]byte, 4)
	binary.LittleEndian.PutUint32(igpIdB, igp.IgpId)
	igpB, err := igp.Marshal()
	if err != nil {
		return err
	}
	store.Set(igpIdB, igpB)
	return nil
}

// createIgp creates the IGP with the given owner and stores it
func (k Keeper) createIgp(ctx sdk.Context, igp_id uint32, owner sdk.AccAddress) error {
	store := k.getIgpStore(ctx)
	igpIdB := make([]byte, 4)
	binary.LittleEndian.PutUint32(igpIdB, igp_id)
	if store.Has(igpIdB) {
		return types.ErrInvalidIgp.Wrapf("igp %d already exists and cannot be created", igp_id)
	}

	newIgp := types.Igp{
		Owner: owner.String(),
		IgpId: igp_id,
	}

	igp, err := newIgp.Marshal()
	if err != nil {
		return err
	}
	store.Set(igpIdB, igp)
	return nil
}
