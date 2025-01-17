package ml

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/terra-money/core/v2/testutil/sample"
	mlsimulation "github.com/terra-money/core/v2/x/ml/simulation"
	"github.com/terra-money/core/v2/x/ml/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = mlsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgGenerate = "op_weight_msg_generate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGenerate int = 100

	opWeightMsgConditional = "op_weight_msg_conditional"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConditional int = 100

	opWeightMsgLlm = "op_weight_msg_llm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLlm int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	mlGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&mlGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgGenerate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGenerate, &weightMsgGenerate, nil,
		func(_ *rand.Rand) {
			weightMsgGenerate = defaultWeightMsgGenerate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGenerate,
		mlsimulation.SimulateMsgGenerate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConditional int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgConditional, &weightMsgConditional, nil,
		func(_ *rand.Rand) {
			weightMsgConditional = defaultWeightMsgConditional
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConditional,
		mlsimulation.SimulateMsgConditional(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLlm int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLlm, &weightMsgLlm, nil,
		func(_ *rand.Rand) {
			weightMsgLlm = defaultWeightMsgLlm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLlm,
		mlsimulation.SimulateMsgLlm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
