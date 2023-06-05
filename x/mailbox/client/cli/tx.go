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
		Short:   "Dispatch message",
		Long:    "Dispatch a message via hyperlane",
		Example: fmt.Sprintf("%s tx %s dispatch <destination-domain> <recipient-address> <message-body>", version.AppName, types.ModuleName),
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
				Sender:            clientCtx.GetFromAddress().String(),
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
		Use:     "process <metadata> <message>",
		Short:   "Process message",
		Long:    "Process a message via hyperlane",
		Example: fmt.Sprintf("%s tx %s process <metadata> <message>", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			metadata := args[0]
			message := args[1]

			msg := &types.MsgProcess{
				Sender:   clientCtx.GetFromAddress().String(),
				Metadata: metadata,
				Message:  message,
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
