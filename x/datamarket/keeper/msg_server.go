package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (server msgServer) UploadData(goCtx context.Context, msg *types.MsgUploadData) (*types.MsgUploadDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := server.Keeper.UploadData(ctx, msg.Class, msg.Uploader, msg.Url)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgUploadData,
			sdk.NewAttribute(types.AttributeClass, msg.Class),
			sdk.NewAttribute(types.AttributeUploader, msg.Uploader),
			sdk.NewAttribute(types.AttributeUrl, msg.Url),
		),
	})
	return &types.MsgUploadDataResponse{}, nil
}

func (server msgServer) BuyData(goCtx context.Context, msg *types.MsgBuyData) (*types.MsgBuyDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := server.Keeper.BuyData(ctx, msg.Buyer, msg.Class)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgBuyData,
			sdk.NewAttribute(types.AttributeBuyer, msg.Buyer),
			sdk.NewAttribute(types.AttributeClass, msg.Class),
		),
	})
	dataSet, _ := server.Keeper.getDataByKey(ctx, types.DataSetKey, msg.Class)
	return &types.MsgBuyDataResponse{DataSet: &dataSet}, nil
}

func (server msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if server.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", server.authority, msg.Authority)
	}

	err := server.Keeper.SetParams(ctx, msg.Params)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeMsgUpdateParams,
			sdk.NewAttribute(types.AttributeDataPrice, msg.Params.DataPrice.String()),
			sdk.NewAttribute(types.AttributeFeePercentage, fmt.Sprintf("%f", msg.Params.FeePercentage)),
		),
	})
	return &types.MsgUpdateParamsResponse{}, nil
}

func (server msgServer) MintTo(goCtx context.Context, msg *types.MsgMintTo) (*types.MsgMintToResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

}
