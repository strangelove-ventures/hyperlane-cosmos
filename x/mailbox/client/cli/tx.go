package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func dispatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dispatch <destination-domain> <recipient-address> <message-body>",
		Short:   "Set default ISM",
		Long:    "Sets the default ISM for the mailbox",
		Example: fmt.Sprintf("%s tx %s set-default-ism [ism-hash]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			destinationDomainRaw, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return err
			}

			destinationDomain := uint32(destinationDomainRaw)

			recipientAddress := args[1]
			messageBody := args[2]

			msg := &types.MsgDispatch{
				DestinationDomain: destinationDomain,
				RecipientAddress:  recipientAddress,
				MessageBody:       messageBody,
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

func processCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-default-ism [ism-hash]",
		Short:   "Set default ISM",
		Long:    "Sets the default ISM for the mailbox",
		Example: fmt.Sprintf("%s tx %s set-default-ism [ism-hash]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ismHash := args[0]

			msg := &types.MsgSetDefaultIsm{
				IsmHash: ismHash,
				Signer:  clientCtx.GetFromAddress().String(),
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

// newStoreCodeCmd returns the command to create a MsgStoreCode transaction
func setDefaultIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-default-ism [ism-hash]",
		Short:   "Set default ISM",
		Long:    "Sets the default ISM for the mailbox",
		Example: fmt.Sprintf("%s tx %s set-default-ism [ism-hash]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ismHash := args[0]

			msg := &types.MsgSetDefaultIsm{
				IsmHash: ismHash,
				Signer:  clientCtx.GetFromAddress().String(),
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
