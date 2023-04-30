package keeper

import (
	"github.com/terra-money/core/v2/x/ml/types"
)

var _ types.QueryServer = Keeper{}
