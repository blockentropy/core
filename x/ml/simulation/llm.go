package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/terra-money/core/v2/x/ml/keeper"
	"github.com/terra-money/core/v2/x/ml/types"
)

func SimulateMsgLlm(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLlm{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Llm simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Llm simulation not implemented"), nil, nil
	}
}
