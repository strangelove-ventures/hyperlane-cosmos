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
)

// newStoreCodeCmd returns the command to create a MsgStoreCode transaction
func setDefaultIsmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-default-ism [path/to/ism.json]",
		Short:   "Set default ISM",
		Long:    strings.TrimSpace(
			fmt.Sprintf(`Sets the default ISM for the mailbox.
The ISM should be defined in a JSON file.
Example:
$ %s tx %s set-default-ism path/to/ism.json

Where ism.json contains:

{
	"validator_pub_keys": [
		"0x123456789",
		"0x234567890",
		"0x345678901"
	]
	"threshold": 2
}			`, version.AppName, types.ModuleName)),
		Example: fmt.Sprintf("%s tx %s set-default-ism [ism-hash]", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			ism, err := parseIsmFile(args[0])
			if err != nil {
				return err
			}

			msg := &types.MsgSetDefaultIsm{
				ValidatorPubKeys: ism.ValidatorPubKeys,
				Threshold: ism.Threshold,
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

type ism struct {
	ValidatorPubKeys [][]byte `json:"validator_pub_keys"`
	Threshold uint32 `json:"threshold"`
}

// parseIsmFile reads and parses the ISM file
func parseIsmFile(path string) (ism, error) {
	var ism ism

	contents, err := os.ReadFile(path)
	if err != nil {
		return ism, fmt.Errorf("invalid path")
	}

	err = json.Unmarshal(contents, &ism)
	if err != nil {
		return ism, fmt.Errorf("cannot unmarshal json file")
	}

	return ism, nil
}