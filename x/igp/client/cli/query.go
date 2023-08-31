package cli

import (
	"context"
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

// quoteGasPaymentCmd Get the amount and denomination (cost) to pay for message delivery
func quoteGasPaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "quoteGasPayment <igp_id> <destination_domain> <gas_amount>",
		Short:   "quoteGasPayment",
		Long:    "quoteGasPayment - get the expected cost for relay of a message to a remote domain",
		Example: fmt.Sprintf("%s query %s tree", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			igpId, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			destinationDomain, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return err
			}

			gasAmt, ok := math.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("%s is not a valid exchange rate (must be Integer)", args[2])
			}

			req := &types.QuoteGasPaymentRequest{IgpId: uint32(igpId), DestinationDomain: uint32(destinationDomain), GasAmount: gasAmt}
			res, err := queryClient.QuoteGasPayment(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
