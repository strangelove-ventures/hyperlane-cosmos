package keeper

import (
	"context"
	"errors"
	"strconv"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

var _ types.MsgServer = (*Keeper)(nil)

// NewMsgServerImpl return an implementation of the igp MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return keeper
}

// PayForGas Make payments for relayer to deliver message to a destination domain
func (k Keeper) PayForGas(goCtx context.Context, msg *types.MsgPayForGas) (*types.MsgPayForGasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	igp, err := k.getIgp(ctx, msg.IgpId)
	if err != nil {
		return nil, err
	}

	beneficiary := igp.GetBeneficiary()
	if beneficiary == "" {
		beneficiary = igp.GetOwner()
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	beneficiaryAcc, err := sdk.AccAddressFromBech32(beneficiary)
	if err != nil {
		return nil, err
	}

	// Get the expected payment amount and denomination
	quoteGasResp, err := k.QuoteGasPayment(ctx, &types.QuoteGasPaymentRequest{IgpId: msg.IgpId, DestinationDomain: msg.DestinationDomain, GasAmount: msg.GasAmount})
	if err != nil {
		return nil, err
	}

	// Verify that the payment is in the chain's native denom
	if quoteGasResp.Denom != msg.MaximumPayment.Denom {
		return nil, types.ErrInvalidPaymentDenom.Wrapf("Payment provided in denom '%s' but require in denom %s", msg.MaximumPayment.Denom, quoteGasResp.Denom)
	}

	requiredPayment := quoteGasResp.Amount
	if msg.MaximumPayment.Amount.LT(requiredPayment) {
		return nil, errors.New("insufficient payment")
	}

	store := k.getGasPaidStore(ctx, msg.DestinationDomain, beneficiaryAcc)

	// message gas is already paid for, deny another payment
	if store.Has([]byte(msg.MessageId)) {
		return nil, types.ErrGasPaid.Wrapf("Message %s gas already paid to domain %d", msg.MessageId, msg.DestinationDomain)
	}

	gasPayment := sdk.NewCoin(quoteGasResp.Denom, requiredPayment)

	// This implementation does not require that beneficiaries Claim() payments.
	// The payment is sent directly to the beneficiary (not escrowed).
	err = k.sendKeeper.SendCoins(ctx, sender, beneficiaryAcc, sdk.NewCoins(gasPayment))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePayForGas,
			sdk.NewAttribute(types.AttributeMessageId, msg.MessageId),
			sdk.NewAttribute(types.AttributeGasAmount, msg.GasAmount.String()),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeBeneficiary, beneficiary),
			sdk.NewAttribute(types.AttributePayment, gasPayment.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	store.Set([]byte(msg.MessageId), []byte(gasPayment.String()))
	return &types.MsgPayForGasResponse{}, nil
}

func (k Keeper) SetRemoteGasData(goCtx context.Context, msg *types.MsgSetRemoteGasData) (*types.MsgSetRemoteGasDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	igp, err := k.getIgp(ctx, msg.IgpId)
	if err != nil {
		return nil, err
	}

	oracle, ok := igp.Oracles[msg.RemoteDomain]

	// The oracle can only be updated if the msg.Sender owns the oracle.
	if !ok {
		return nil, types.ErrOracleUnauthorized.Wrapf("oracle with destination %d does not exist for IGP %d", msg.RemoteDomain, igp.IgpId)
	} else if oracle.GasOracle != msg.Sender {
		return nil, types.ErrOracleUnauthorized.Wrapf("account %s is unauthorized to configure oracle for IGP %d and remote domain %d", msg.Sender, igp.IgpId, msg.RemoteDomain)
	}

	// Store the updated exchange rate and gas price for the oracle
	oracle.GasPrice = msg.GasPrice
	oracle.TokenExchangeRate = msg.TokenExchangeRate
	err = k.setIgp(ctx, igp)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeGasDataSet,
			sdk.NewAttribute(types.AttributeRemoteDomain, strconv.FormatUint(uint64(msg.RemoteDomain), 10)),
			sdk.NewAttribute(types.AttributeTokenExchangeRate, msg.TokenExchangeRate.String()),
			sdk.NewAttribute(types.AttributeGasPrice, msg.GasPrice.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgSetRemoteGasDataResponse{}, nil
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
	err = k.setIgp(ctx, igp)
	if err != nil {
		return nil, err
	}

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
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Configs) == 0 {
		return nil, errors.New("invalid GasOracleConfig")
	}

	igps := map[uint32]*types.Igp{}

	events := sdk.Events{}

	for _, oracleConfig := range msg.Configs {
		var igp *types.Igp

		// Lookup the IGP
		igp, ok := igps[oracleConfig.IgpId]
		if !ok {
			igp, err = k.getIgp(ctx, oracleConfig.IgpId)
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

			events = events.AppendEvent(sdk.NewEvent(
				types.EventTypeCreateOracle,
				sdk.NewAttribute(types.AttributeOracleAddress, oracleConfig.GasOracle),
			))
		} else {
			// This is an existing oracle, confirm authorization to update it.
			// The oracle can be updated if the msg.Sender owns the IGP or the oracle itself.
			if igp.Owner != msg.Sender && oracle.GasOracle != msg.Sender {
				return nil, types.ErrOracleUnauthorized.Wrapf("account %s is unauthorized to configure existing oracle for IGP %d with owner %s", msg.Sender, igp.IgpId, igp.Owner)
			}

			events = events.AppendEvent(sdk.NewEvent(
				types.EventTypeUpdateOracle,
				sdk.NewAttribute(types.AttributeOracleAddress, oracleConfig.GasOracle),
			))
		}

		// Configure the address that can update the gas oracle config
		oracle.GasOracle = oracleConfig.GasOracle
		err = k.setIgp(ctx, igp)
		if err != nil {
			return nil, err
		}
	}

	events = events.AppendEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
	))
	ctx.EventManager().EmitEvents(events)

	return &types.MsgSetGasOraclesResponse{}, nil
}

func (k Keeper) CreateIgp(goCtx context.Context, msg *types.MsgCreateIgp) (*types.MsgCreateIgpResponse, error) {
	validExchRate := msg.TokenExchangeRateScale.IsZero() || msg.TokenExchangeRateScale.GTE(math.OneInt())
	if !validExchRate {
		return nil, types.ErrExchangeRateScale.Wrapf("provided %s, exchange rate should be power of ten", msg.TokenExchangeRateScale.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	newIgp := types.Igp{
		Owner:                  msg.Sender,
		Beneficiary:            msg.Beneficiary,
		TokenExchangeRateScale: math.NewInt(1e10),
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
	err := k.setIgp(ctx, &newIgp)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateIgp,
			sdk.NewAttribute(types.AttributeIgpId, strconv.FormatUint(uint64(igp_id), 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

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
	err = k.setIgp(ctx, igp)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetBeneficiary,
			sdk.NewAttribute(types.AttributeIgpId, strconv.FormatUint(uint64(msg.IgpId), 10)),
			sdk.NewAttribute(types.AttributeBeneficiary, msg.Address),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgSetBeneficiaryResponse{}, nil
}
