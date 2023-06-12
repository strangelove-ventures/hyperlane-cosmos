package keeper

import (
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
	domain      uint32
	Tree        *imt.Tree
	Delivered   map[string]bool
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, cwKeeper *cosmwasm.Keeper, ismKeeper *ismkeeper.Keeper, domain uint32) Keeper {
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
		domain:      domain,
		pcwKeeper:   cosmwasm.NewDefaultPermissionKeeper(cwKeeper),
		Tree:        &imt.Tree{},
		Delivered:   map[string]bool{},
	}
}

func (k Keeper) VerifyMessage(messageBytes []byte) (string, error) {
	if common.Version(messageBytes) != k.version {
		return "", types.ErrMsgInvalidVersion
	}

	if common.Destination(messageBytes) != k.domain {
		return "", types.ErrMsgInvalidDomain
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
