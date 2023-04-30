package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/terra-money/core/v2/testutil/keeper"
	"github.com/terra-money/core/v2/x/ml/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MlKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
