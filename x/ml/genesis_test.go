package ml_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/terra-money/core/v2/testutil/keeper"
	"github.com/terra-money/core/v2/testutil/nullify"
	"github.com/terra-money/core/v2/x/ml"
	"github.com/terra-money/core/v2/x/ml/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MlKeeper(t)
	ml.InitGenesis(ctx, *k, genesisState)
	got := ml.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
