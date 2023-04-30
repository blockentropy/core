package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/v2/x/ml/types"
)

func (k msgServer) Generate(goCtx context.Context, msg *types.MsgGenerate) (*types.MsgGenerateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGenerateResponse{}, nil
}
