package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

// getOriginsDefaultIsmCmd defines the command to query the default ISM
func getOriginsDefaultIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "default-ism [origin]",
		Short:   "Query default ISM for origin",
		Long:    "Query default ISM for a specific origin",
		Example: fmt.Sprintf("%s query %s default-ism [origin]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			origin, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			req := types.QueryOriginsDefaultIsmRequest{
				Origin: uint32(origin),
			}

			res, err := queryClient.OriginsDefaultIsm(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// getAllDefaultIsmCmd defines the command to query the default ISM
func getAllDefaultIsmsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "all-default-isms",
		Short:   "Query all default ISMs",
		Long:    "Query all default ISMs",
		Example: fmt.Sprintf("%s query %s all-default-isms", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryAllDefaultIsmsRequest{}

			res, err := queryClient.AllDefaultIsms(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// getCustomIsmCmd defines the command to query a custom ISM
func getCustomIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "custom-ism [ism-id]",
		Short:   "Query custom ISM",
		Long:    "Query custom ISM given an ism-id",
		Example: fmt.Sprintf("%s query %s custom-ism [ism-id]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			ismId, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			req := types.QueryCustomIsmRequest{
				IsmId: uint32(ismId),
			}

			res, err := queryClient.CustomIsm(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// getAllCustomIsmsCmd defines the command to query all custom ISMs
func getAllCustomIsmsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "all-custom-isms",
		Short:   "Query all custom ISMs",
		Long:    "Query all custom ISMs",
		Example: fmt.Sprintf("%s query %s all-custom-isms", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryAllCustomIsmsRequest{}

			res, err := queryClient.AllCustomIsms(context.Background(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
