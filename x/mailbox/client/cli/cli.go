package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

// GetQueryCmd returns the query commands for IBC channels
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Hyperlane mailbox module query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		getCurrentTreeMetadataCmd(),
		getDomain(),
	)

	return queryCmd
}

// NewTxCmd returns a CLI command handler for all x/ibc channel transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Hyperlane mailbox module transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		dispatchCmd(),
		processCmd(),
	)

	return txCmd
}
