package cli

import (
	"errors"
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

func createIgpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "createigp <beneficiary-address>",
		Short:   "Create new IGP",
		Long:    "Create new hyperlane IGP",
		Example: fmt.Sprintf("%s tx %s createigp cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			beneficiary := args[0]

			msg := &types.MsgCreateIgp{
				Sender:      clientCtx.GetFromAddress().String(),
				Beneficiary: beneficiary,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func msgPaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "msgpay <message-id> <destination-domain> <destination-gas-amount> <igp-id> <max-payment>",
		Short:   "Dispatch message",
		Long:    "Dispatch a message via hyperlane",
		Example: fmt.Sprintf("%s tx %s msgpay <destination-domain> <destination-gas-amount> <max-payment> <igp-id>", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// TODO: this is supposed to be a hash so we should try to validate it properly
			messageId := args[0]

			destinationDomainRaw, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			destinationDomain := uint32(destinationDomainRaw)
			destGasAmount, ok := math.NewIntFromString(args[2])
			if !ok {
				return errors.New("destination-gas-amount must be Integer type")
			}

			igpId, err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				return errors.New("Igp ID must be integer")
			}

			var maxPayment sdk.Coin
			if len(args) >= 5 {
				coin, err := sdk.ParseCoinNormalized(args[4])
				if err == nil {
					maxPayment = coin
				}
			}

			msg := &types.MsgPayForGas{
				Sender:            clientCtx.GetFromAddress().String(),
				DestinationDomain: destinationDomain,
				GasAmount:         destGasAmount,
				MessageId:         messageId,
				IgpId:             uint32(igpId),
				MaximumPayment:    maxPayment,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
