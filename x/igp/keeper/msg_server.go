package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the ism MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// PayForGas defines a rpc handler method for MsgPayForGas
// TODO: Refunds need to be implemented, see https://docs.hyperlane.xyz/docs/apis-and-sdks/interchain-gas-paymaster-api#refunds
func (k Keeper) PayForGas(goCtx context.Context, msg *types.MsgPayForGas) (*types.MsgPayForGasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	var relayer sdk.AccAddress
	if msg.RelayerAddress != "" {
		relayer, err = sdk.AccAddressFromBech32(msg.RelayerAddress)
		if err != nil {
			return nil, types.ErrInvalidRelayer.Wrapf("relayer %s is not a valid bech32 address", msg.RelayerAddress)
		}
	} else {
		relayer = k.getDefaultRelayer(ctx)
		if relayer == nil {
			return nil, types.ErrInvalidRelayer.Wrapf("default relayer is not configured. include a relayer in MsgPayForGas.")
		}
	}

	store := k.getGasPaidStore(ctx, msg.DestinationDomain, relayer)

	// message gas is already paid for, deny another payment
	if store.Has([]byte(msg.MessageId)) {
		return nil, types.ErrGasPaid.Wrapf("Message %s gas already paid to domain %d", msg.MessageId, msg.DestinationDomain)
	}

	gasDenom := getGasPaymentDenom()
	amountPayable := quoteGasPayment(msg.DestinationDomain, msg.GasAmount)
	gasPayment := sdk.NewCoin(gasDenom, amountPayable)

	// TODO: IMPORTANT: Technically, funds should be escrowed in PayForGas, then Claim()ed by the relayer.
	// However in Cosmos there is really no reason to Claim(), since nothing happens if a relayer never claims the payment.
	// In other words... sender can never recover the funds either way in the current hyperlane spec.
	k.sendKeeper.SendCoins(ctx, sender, relayer, sdk.NewCoins(gasPayment))

	//TODO: improve error message
	if err != nil {
		return nil, err
	}
	store.Set([]byte(msg.MessageId), []byte(gasPayment.String()))
	return &types.MsgPayForGasResponse{}, nil
}

func (k Keeper) SetDestinationGasOverhead(goCtx context.Context, msg *types.MsgSetDestinationGasOverhead) (*types.MsgSetDestinationGasOverheadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	key := types.GasOverheadKey(msg.DestinationDomain)
	gasOh, err := msg.GasOverhead.Marshal()
	if err != nil {
		return nil, err
	}
	store.Set(key, gasOh)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetGasOverhead,
			sdk.NewAttribute(types.AttributeDestination, strconv.FormatUint(uint64(msg.DestinationDomain), 10)),
			sdk.NewAttribute(types.AttributeOverheadAmount, msg.GasOverhead.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})
	return &types.MsgSetDestinationGasOverheadResponse{}, nil
}

// Claim defines a rpc handler method for MsgClaims
func (k Keeper) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgClaimResponse{}, nil
}

// SetGasOracles defines a rpc handler method for MsgSetGasOracles
func (k Keeper) SetGasOracles(goCtx context.Context, msg *types.MsgSetGasOracles) (*types.MsgSetGasOraclesResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgSetGasOraclesResponse{}, nil
}

// SetGasOracles defines a rpc handler method for MsgSetGasOracles
func (k Keeper) SetBeneficiary(goCtx context.Context, msg *types.MsgSetBeneficiary) (*types.MsgSetBeneficiaryResponse, error) {
	//_ := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgSetBeneficiaryResponse{}, nil
}
