package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

// GetQueryCmd returns the query commands for Hyperlane IGP module commands
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Hyperlane announce module query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// queryCmd.AddCommand()
	return queryCmd
}

// NewTxCmd returns a CLI command handler for all Hyperlane IGP module transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Hyperlane announce module transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
	// msgPaymentCmd(),
	// createIgpCmd(),
	// createOracleCmd(),
	// setGasPriceCmd(),
	)

	return txCmd
}
