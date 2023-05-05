package cli

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

// getDefaultIsmCmd defines the command to query the default ISM
func getDefaultIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "default-ism",
		Short:   "Query default ISM",
		Long:    "Query default ISM",
		Example: fmt.Sprintf("%s query %s default-ism", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryDefaultIsmRequest{}

			res, err := queryClient.DefaultIsm(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// getContractIsmCmd defines the command to query the contract ISM
func getContractIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "contract-ism [contract-addr]",
		Short:   "Query contract ISM",
		Long:    "Query contract ISM",
		Example: fmt.Sprintf("%s query %s contract-ism [contract-addr]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			contractAddrHex := args[0]
			contractAddr, err := hex.DecodeString(contractAddrHex)
			if err != nil {
				return err
			}

			req := types.QueryContractIsmRequest{
				ContractAddr: contractAddr,
			}

			res, err := queryClient.ContractIsm(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
