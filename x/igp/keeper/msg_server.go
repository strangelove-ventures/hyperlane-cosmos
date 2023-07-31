package keeper

import (
	"context"
	"errors"
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

	// Note: This implementation does not require that beneficiaries Claim() payments.
	// Instead, the payment is sent directly to the beneficiary here (not escrowed).
	// TODO: compare payment to quoteGasPayment and ensure we do not overcharge the payer.
	k.sendKeeper.SendCoins(ctx, sender, relayer, sdk.NewCoins(gasPayment))
	if err != nil {
		return nil, err
	}
	store.Set([]byte(msg.MessageId), []byte(gasPayment.String()))
	return &types.MsgPayForGasResponse{}, nil
}

func (k Keeper) SetDestinationGasOverhead(goCtx context.Context, msg *types.MsgSetDestinationGasOverhead) (*types.MsgSetDestinationGasOverheadResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	igp, err := k.getIgp(ctx, msg.IgpId)
	if err != nil {
		return nil, err
	}

	oracle, ok := igp.Oracles[msg.DestinationDomain]

	if !ok {
		return nil, types.ErrOracleUnauthorized.Wrapf("oracle with destination %d does not exist for IGP %d", msg.DestinationDomain, igp.IgpId)
	} else {
		// This is an existing oracle, confirm authorization to update it.
		// The oracle can be updated if the msg.Sender owns the IGP or the oracle itself.
		if igp.Owner != msg.Sender && oracle.GasOracle != msg.Sender {
			return nil, types.ErrOracleUnauthorized.Wrapf("account %s is unauthorized to configure existing oracle for IGP %d with owner %s", msg.Sender, igp.IgpId, igp.Owner)
		}
	}

	// Configure the gas overhead for the oracle
	oracle.GasOverhead = msg.GasOverhead
	k.setIgp(ctx, igp)

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

// SetGasOracles defines a rpc handler method for MsgSetGasOracles
func (k Keeper) SetGasOracles(goCtx context.Context, msg *types.MsgSetGasOracles) (*types.MsgSetGasOraclesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Configs) == 0 {
		return nil, errors.New("invalid GasOracleConfig")
	}

	igps := map[uint32]*types.Igp{}

	for _, oracleConfig := range msg.Configs {
		var igp *types.Igp

		// Lookup the IGP
		igp, ok := igps[oracleConfig.IgpId]
		if !ok {
			igp, err := k.getIgp(ctx, oracleConfig.IgpId)
			if err != nil {
				return nil, err
			}
			igps[oracleConfig.IgpId] = igp
		}

		oracle, existingOracle := igp.Oracles[oracleConfig.RemoteDomain]

		// This is a new oracle, create and set it on the IGP
		if !existingOracle {
			if igp.Owner != msg.Sender {
				return nil, types.ErrOracleUnauthorized.Wrapf("account %s is unauthorized to configure new oracle for IGP %d with owner %s", msg.Sender, igp.IgpId, igp.Owner)
			}

			oracle = &types.GasOracle{}
			igp.Oracles[oracleConfig.RemoteDomain] = oracle
		} else {
			// This is an existing oracle, confirm authorization to update it.
			// The oracle can be updated if the msg.Sender owns the IGP or the oracle itself.
			if igp.Owner != msg.Sender && oracle.GasOracle != msg.Sender {
				return nil, types.ErrOracleUnauthorized.Wrapf("account %s is unauthorized to configure existing oracle for IGP %d with owner %s", msg.Sender, igp.IgpId, igp.Owner)
			}
		}

		// Configure the owner who can update the gas oracle (this can be different than the IGP owner)
		oracle.GasOracle = oracleConfig.GasOracle
		// TODO: set gas prices, overhead (optional)
	}

	return &types.MsgSetGasOraclesResponse{}, nil
}

func (k Keeper) CreateIgp(goCtx context.Context, msg *types.MsgCreateIgp) (*types.MsgCreateIgpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newIgp := types.Igp{
		Owner:       msg.Sender,
		Beneficiary: msg.Beneficiary,
	}

	igp_id := uint32(0)

	for {
		_, err := k.getIgp(ctx, igp_id)
		if err != nil {
			break
		}
		igp_id += 1
	}
	newIgp.IgpId = igp_id
	k.setIgp(ctx, &newIgp)
	return &types.MsgCreateIgpResponse{IgpId: igp_id}, nil
}

// SetBeneficiary updates the IGP's beneficiary (account sent relayer gas payments)
func (k Keeper) SetBeneficiary(goCtx context.Context, msg *types.MsgSetBeneficiary) (*types.MsgSetBeneficiaryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	igp, err := k.getIgp(ctx, msg.IgpId)
	if err != nil {
		return nil, err
	}

	// Only the IGP owner can change the beneficiary
	if igp.Owner != msg.Sender {
		return nil, types.ErrBeneficiaryUnauthorized.Wrapf("account %s is unauthorized to configure beneficiary for IGP %d and owner %s", msg.Sender, igp.IgpId, igp.Owner)
	}

	igp.Beneficiary = msg.Address
	k.setIgp(ctx, igp)

	return &types.MsgSetBeneficiaryResponse{}, nil
}
