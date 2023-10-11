package cli

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func announceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "announce <hex-validator-address> <storageLocation> <hex-validator-signature>",
		Short:   "announce - Announces a validator signature location",
		Example: fmt.Sprintf("%s tx %s announce 0xFFFFFFFFFFFFFFFFFFFF", version.AppName, types.ModuleName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valHex := args[0]
			storageLocation := args[1]
			sigHex := args[2]

			if strings.Contains(valHex, "0x") {
				_, valHex, _ = strings.Cut(valHex, "0x") // remove any 0x prefixes so we can properly decode
			}
			if strings.Contains(sigHex, "0x") {
				_, sigHex, _ = strings.Cut(sigHex, "0x") // remove any 0x prefixes so we can properly decode
			}

			address, err := hex.DecodeString(valHex)
			if err != nil {
				return err
			}
			sig, err := hex.DecodeString(sigHex)
			if err != nil {
				return err
			}

			msg := &types.MsgAnnouncement{
				Sender:          clientCtx.GetFromAddress().String(),
				Validator:       address,
				Signature:       sig,
				StorageLocation: storageLocation,
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
