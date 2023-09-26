package keeper

import (
	"context"
	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (server msgServer) UploadData(goCtx context.Context, msg *types.MsgUploadData) (*types.MsgUploadDataResponse, error) {

}

func (server msgServer) BuyData(goCtx context.Context, msg *types.MsgBuyData) (*types.MsgBuyDataResponse, error) {

}

func (server msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {

}
