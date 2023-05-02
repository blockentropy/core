package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/v2/x/ml/types"
)

func (k msgServer) Llm(goCtx context.Context, msg *types.MsgLlm) (*types.MsgLlmResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLlmResponse{}, nil
}
