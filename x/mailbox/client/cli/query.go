package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

// getCurrentTreeMetadataCmd defines the command to query the current tree metadata
func getCurrentTreeMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tree-metadata",
		Short:   "Query tree-metadata",
		Long:    "Query current tree metadata",
		Example: fmt.Sprintf("%s query %s tree-metadata", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryCurrentTreeMetadataRequest{}

			res, err := queryClient.CurrentTreeMetadata(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// getCurrentTreeCmd defines the command to query the current tree
func getCurrentTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tree",
		Short:   "Query tree",
		Long:    "Query current tree",
		Example: fmt.Sprintf("%s query %s tree", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryCurrentTreeRequest{}

			res, err := queryClient.CurrentTree(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// getCurrentTreeMetadataCmd defines the command to query the current tree metadata
func getDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "domain",
		Short:   "Query domain",
		Long:    "Query domain",
		Example: fmt.Sprintf("%s query %s domain", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryDomainRequest{}

			res, err := queryClient.Domain(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
