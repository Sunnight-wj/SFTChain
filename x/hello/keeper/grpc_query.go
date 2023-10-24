package keeper

import (
	"context"
	"github.com/CosmosContracts/juno/v17/x/hello/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)

	return &types.QueryParamsResponse{Params: params}, nil
}
