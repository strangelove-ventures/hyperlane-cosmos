package keeper

import (
	"context"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	cosmwasm "github.com/CosmWasm/wasmd/x/wasm/keeper"

	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	ismkeeper "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer
	cwKeeper  *cosmwasm.Keeper
	pcwKeeper *cosmwasm.PermissionedKeeper
	ismKeeper *ismkeeper.Keeper

	storeKey    storetypes.StoreKey
	cdc         codec.BinaryCodec
	authority   string
	mailboxAddr sdk.AccAddress
	version     byte

	Tree      *imt.Tree
	Delivered map[string]bool
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, cwKeeper *cosmwasm.Keeper, ismKeeper *ismkeeper.Keeper) Keeper {
	// governance authority
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	return Keeper{
		cdc:         cdc,
		storeKey:    key,
		cwKeeper:    cwKeeper,
		ismKeeper:   ismKeeper,
		authority:   authority.String(),
		mailboxAddr: authtypes.NewModuleAddress(types.ModuleName),
		version:     0,
		pcwKeeper:   cosmwasm.NewDefaultPermissionKeeper(cwKeeper),
		Tree:        &imt.Tree{},
		Delivered:   map[string]bool{},
	}
}

func (k *Keeper) SetDomain(c context.Context, domain uint32) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	res := make([]byte, 4)
	binary.LittleEndian.PutUint32(res, domain)
	store.Set(types.DomainKey, res)
}

func (k Keeper) VerifyMessage(c context.Context, messageBytes []byte) (string, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.DomainKey)
	domain := binary.LittleEndian.Uint32(b)

	if common.Version(messageBytes) != k.version {
		return "", types.ErrMsgInvalidVersion
	}

	destGiven := common.Destination(messageBytes)
	if common.Destination(messageBytes) != domain {
		return "", types.ErrMsgInvalidDomain.Wrapf("Message destination %d is invalid. Acceptable domain is %d", destGiven, domain)
	}

	idBytes := common.Id(messageBytes)
	id := hexutil.Encode(idBytes)

	// Verify message has not been delivered yet
	val, ok := k.Delivered[id]
	if ok && val {
		return "", types.ErrMsgDelivered
	}

	return id, nil
}
