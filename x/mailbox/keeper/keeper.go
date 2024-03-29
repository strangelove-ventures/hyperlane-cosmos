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
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/receiver"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

var _ ReadOnlyMailboxKeeper = (*Keeper)(nil)

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

	receivers []receiver.Receiver
}

type ReadOnlyMailboxKeeper interface {
	GetMailboxAddress() []byte
	GetDomain(context.Context) uint32
}

// Get 32 byte mailbox address, pad if necessary
func (k Keeper) GetMailboxAddress() []byte {
	mailboxAddr := k.mailboxAddr
	for len(mailboxAddr) < 32 {
		padding := make([]byte, 1)
		mailboxAddr = append(padding, mailboxAddr...)
	}
	return mailboxAddr
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, ismKeeper *ismkeeper.Keeper) Keeper {
	// governance authority
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	return Keeper{
		cdc:         cdc,
		storeKey:    key,
		ismKeeper:   ismKeeper,
		authority:   authority.String(),
		mailboxAddr: authtypes.NewModuleAddress(types.ModuleName),
		version:     0,
		pcwKeeper:   nil,
		receivers:   []receiver.Receiver{},
	}
}

func (k *Keeper) AddReceiver(receiver *receiver.Receiver) {
	k.receivers = append(k.receivers, *receiver)
}

func (k *Keeper) AddWasmReceiver(cwKeeper *cosmwasm.Keeper) {
	k.cwKeeper = cwKeeper
	k.pcwKeeper = cosmwasm.NewDefaultPermissionKeeper(cwKeeper)
}

func (k *Keeper) SetDomain(c context.Context, domain uint32) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	res := make([]byte, 4)
	binary.LittleEndian.PutUint32(res, domain)
	store.Set(types.DomainKey, res)
}

func (k Keeper) GetDomain(c context.Context) uint32 {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	// Get the domain from the store
	res := store.Get(types.DomainKey)
	if res == nil {
		return 0
	}

	domain := binary.LittleEndian.Uint32(res)

	return domain
}

// Stores the proto type tree
func (k *Keeper) StoreTree(c context.Context, tree types.Tree) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	treeBz := k.cdc.MustMarshal(&tree)
	store.Set(types.MailboxIMTKey(), treeBz)
}

// Stores the IMT type tree
func (k *Keeper) StoreImtTree(c context.Context, tree *imt.Tree) {
	protoTree := types.Tree{
		Branch: tree.Branch[:],
		Count:  tree.Count(),
	}
	k.StoreTree(c, protoTree)
}

// Gets the proto type tree
func (k *Keeper) GetTree(c context.Context) types.Tree {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	treeBz := store.Get(types.MailboxIMTKey())
	var tree types.Tree
	k.cdc.MustUnmarshal(treeBz, &tree)

	return tree
}

// Gets the IMT type tree
func (k *Keeper) GetImtTree(c context.Context) *imt.Tree {
	protoTree := k.GetTree(c)
	return imt.InitializeTree(protoTree.Branch, protoTree.Count)
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
	if store.Has(types.MailboxDeliveredKey(id)) {
		return "", types.ErrMsgDelivered
	}

	return id, nil
}
