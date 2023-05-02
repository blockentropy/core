package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/terra-money/core/v2/x/ml/types"
)

var _ = strconv.Itoa(0)

func CmdGenerate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [modality] [model] [prompt] [negprompt] [seed] [machine] [endpoint]",
		Short: "Broadcast message generate",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argModality := args[0]
			argModel := args[1]
			argPrompt := args[2]
			argNegprompt := args[3]
			argSeed := args[4]
			argMachine := args[5]
			argEndpoint := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgGenerate(
				clientCtx.GetFromAddress().String(),
				argModality,
				argModel,
				argPrompt,
				argNegprompt,
				argSeed,
				argMachine,
				argEndpoint,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
