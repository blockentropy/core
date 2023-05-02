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

func CmdConditional() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "conditional [modality] [model] [ctrlmodel] [ctrlinput] [prompt] [negprompt] [seed] [machine]",
		Short: "Broadcast message conditional",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argModality := args[0]
			argModel := args[1]
			argCtrlmodel := args[2]
			argCtrlinput := args[3]
			argPrompt := args[4]
			argNegprompt := args[5]
			argSeed := args[6]
			argMachine := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgConditional(
				clientCtx.GetFromAddress().String(),
				argModality,
				argModel,
				argCtrlmodel,
				argCtrlinput,
				argPrompt,
				argNegprompt,
				argSeed,
				argMachine,
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
