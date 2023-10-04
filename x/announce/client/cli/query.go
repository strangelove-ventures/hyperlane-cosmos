package cli

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

// getAnnouncedStorageLocations Returns a list of all announced storage locations
func getAnnouncedStorageLocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "getAnnouncedStorageLocations <validator_address_hex_csv>",
		Short:   "getAnnouncedStorageLocations",
		Long:    "getAnnouncedStorageLocations - get announced storage locations given a comma separated list of hex validator addresses",
		Example: fmt.Sprintf("%s query %s tree", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			addresses := strings.Split(args[0], ",")
			validators := [][]byte{}
			for _, hexAddr := range addresses {
				_, hexAddr, _ = strings.Cut(hexAddr, "0x") // remove any 0x prefixes so we can properly decode
				address, err := hex.DecodeString(hexAddr)
				if err != nil {
					return fmt.Errorf("Address %s not provided in hex format", hexAddr)
				}
				validators = append(validators, address)
			}

			req := &types.GetAnnouncedStorageLocationsRequest{Validator: validators}
			res, err := queryClient.GetAnnouncedStorageLocations(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// getAnnouncedValidators Returns a list of all announced validators
func getAnnouncedValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "getAnnouncedValidators",
		Short:   "getAnnouncedValidators",
		Long:    "getAnnouncedValidators - get a list of all announced validators",
		Example: fmt.Sprintf("%s query %s tree", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.GetAnnouncedValidatorsRequest{}
			res, err := queryClient.GetAnnouncedValidators(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
