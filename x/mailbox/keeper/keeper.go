package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	cosmwasm "github.com/CosmWasm/wasmd/x/wasm/keeper"
	cwtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/imt"
	ism "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer
	cwKeeper  *cosmwasm.Keeper
	pcwKeeper *cosmwasm.PermissionedKeeper

	storeKey  storetypes.StoreKey
	cdc       codec.BinaryCodec
	authority string
	version   byte
	domain    uint32
	tree      imt.Tree
	delivered map[string]bool
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, cwKeeper *cosmwasm.Keeper, domain uint32) Keeper {
	// governance authority
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	return Keeper{
		cdc:       cdc,
		storeKey:  key,
		cwKeeper:  cwKeeper,
		authority: authority.String(),
		version:   0,
		domain:    domain,
		pcwKeeper: cosmwasm.NewDefaultPermissionKeeper(cwKeeper),
	}
}

func (k Keeper) getRecipientISM() ism.Keeper {
	panic("Implement Me")
}

func (k Keeper) HandleMessage(goCtx context.Context, origin uint32, sender, recipient, body string) ([]byte, error) {

	senderAddr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "sender")
	}

	contractAddr, err := sdk.AccAddressFromBech32(recipient)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "contract")
	}

	var msg cwtypes.RawContractMessage
	err = msg.UnmarshalJSON([]byte(body))
	if err != nil {
		return nil, err
	}

	err = msg.ValidateBasic()
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	data, err := k.pcwKeeper.Execute(ctx, contractAddr, senderAddr, msg.Bytes(), sdk.NewCoins())
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, sender),
	))

	return data, nil
}
