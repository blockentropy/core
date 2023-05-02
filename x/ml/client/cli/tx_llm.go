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

func CmdLlm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "llm [modality] [model] [prompt] [context] [machine] [endpoint]",
		Short: "Broadcast message llm",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argModality := args[0]
			argModel := args[1]
			argPrompt := args[2]
			argContext := args[3]
			argMachine := args[4]
			argEndpoint := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLlm(
				clientCtx.GetFromAddress().String(),
				argModality,
				argModel,
				argPrompt,
				argContext,
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
