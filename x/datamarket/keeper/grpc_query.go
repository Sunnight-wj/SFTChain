package keeper

import (
	"context"
	"fmt"
	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) Data(ctx context.Context, request *types.QueryDataRequest) (*types.QueryDataResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	class := request.Class
	urls, err := k.getDataByKey(sdkCtx, types.DataSetKey, class)
	if err != nil {
		return nil, err
	}

	return &types.QueryDataResponse{DataSet: urls}, nil
}

func (k Keeper) Share(ctx context.Context, request *types.QueryShareRequest) (*types.QueryShareResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	uploaderUrls, _ := k.getDataByKey(sdkCtx, request.Uploader, request.Class)
	allUrls, err := k.getDataByKey(sdkCtx, types.DataSetKey, request.Class)
	if err != nil {
		return nil, err
	}
	if len(allUrls.Urls) == 0 {
		return nil, fmt.Errorf("data set is empty, class: %s", request.Class)
	}
	share := float64(len(uploaderUrls.Urls)) / float64(len(allUrls.Urls))
	return &types.QueryShareResponse{Share: share}, nil
}

func (k Keeper) DataSetFromBuyer(ctx context.Context, request *types.QueryDataSetFromBuyerRequest) (*types.QueryDataSetFromBuyerResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	dataSet, _ := k.getDataSetByBuyer(sdkCtx, request.Buyer)

	return &types.QueryDataSetFromBuyerResponse{Classes: dataSet}, nil
}

func (k Keeper) VipInfo(ctx context.Context, request *types.QueryVipInfoRequest) (*types.QueryVipInfoResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	vipInfo, _ := k.getVipInfo(sdkCtx)

	return &types.QueryVipInfoResponse{VipInfo: vipInfo}, nil
}
