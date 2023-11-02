package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/message_id_multisig"
)

// setDefaultIsmCmd returns the command to set default ISM(s)
func setDefaultIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-default-ism [path/to/ism.json]",
		Short: "Set default ISM for origin",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Sets the default ISM for the mailbox.
The ISM should be defined in a JSON file.
Example:
$ %s tx %s set-default-ism [path/to/ism.json]

Where ism.json contains:
{
	"@type":""/hyperlane.ism.v1.MsgSetDefaultIsm",
	"isms": 
	[
		{
			"origin": 1,
			"abstract_ism": {
				"@type":"/hyperlane.ism.v1.LegacyMultiSig",
				"validator_pub_keys": [
					"0x123456789",
					"0x234567890",
					"0x345678901"
				],
				"threshold": 2
			}
		},
		{
			"origin": 2,
			"abstract_ism": {
				"@type":"/hyperlane.ism.v1.MerkleRootMultiSig",
				"validator_pub_keys": [
					"0x123456789",
					"0x234567890",
					"0x345678901"
				],
				"threshold": 2
			}
		},
		...
	]
}			`, version.AppName, types.ModuleName)),
		Example: fmt.Sprintf("%s tx %s set-default-ism [path/to/ism.json]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contents, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var msg types.MsgSetDefaultIsm
			err = json.Unmarshal(contents, &msg)
			if err != nil {
				return err
			}

			msg.Signer = clientCtx.GetFromAddress().String()

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// createIsmCmd returns the command to create a MsgCreateIsm transaction
func createMultiSigIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-multisig-ism [type] [ism-json-string]",
		Short: "Create a new multisig ISM",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new LegacyMultiSig ISM.
Example:
$ %s tx %s create-ism LegacyMultiSig 
{"validator_pub_keys":["0x123456789...","0x234567890...","0x345678901..."],"threshold":2}`, version.AppName, types.ModuleName)),
		Example: fmt.Sprintf("%s tx %s create-ism [type] [ism-json-string]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			contents := []byte(args[1])

			var ism types.AbstractIsm
			switch args[0] {
			case "LegacyMultiSig":
				var lms legacy_multisig.LegacyMultiSig
				err = json.Unmarshal(contents, &lms)
				if err != nil {
					return err
				}
				ism = &lms
			case "MerkleRootMultiSig":
				var mrms merkle_root_multisig.MerkleRootMultiSig
				err = json.Unmarshal(contents, &mrms)
				if err != nil {
					return err
				}
				ism = &mrms
			case "MessageIdMultiSig":
				var mims message_id_multisig.MessageIdMultiSig
				err = json.Unmarshal(contents, &mims)
				if err != nil {
					return err
				}
				ism = &mims
			}

			ismAny, err := types.PackAbstractIsm(ism)
			if err != nil {
				return err
			}

			msg := types.MsgCreateIsm{
				Signer: clientCtx.GetFromAddress().String(),
				Ism:    ismAny,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
