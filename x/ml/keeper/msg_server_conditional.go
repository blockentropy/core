package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/v2/x/ml/types"
)

func (k msgServer) Conditional(goCtx context.Context, msg *types.MsgConditional) (*types.MsgConditionalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgConditionalResponse{}, nil
}
