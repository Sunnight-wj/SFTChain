package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

func (k Keeper) BuyData(ctx sdk.Context, buyer, class string) error {
	// 付费

	// 写数据库
}

func (k Keeper) chargeForBuyData(ctx sdk.Context, buyer, class string) error {
	cost, err := k.calculateCost(ctx, class)
	if err != nil {
		return err
	}
	buyerAddr, err := sdk.AccAddressFromBech32(buyer)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyerAddr, types.ModuleName, []sdk.Coin{cost})
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) distributeCost(ctx sdk.Context, class string, cost sdk.Coin) error {
	feePercentage := k.GetParams(ctx).FeePercentage
	costAmount := cost.Amount.Uint64()
	feeAmount := uint64(float64(costAmount) * feePercentage)
	rewardAmount := costAmount - feeAmount

}

func (k Keeper) calculateCost(ctx sdk.Context, class string) (sdk.Coin, error) {
	params := k.GetParams(ctx)
	dataSet, _ := k.getDataByKey(ctx, types.DataSetKey, class)
	if len(dataSet.Urls) == 0 {
		return sdk.Coin{}, fmt.Errorf("dataset is empty, class: %s", class)
	}
	dataAmount := int64(len(dataSet.Urls))
	feeAmount := params.DataPrice.Amount.MulRaw(dataAmount)
	fee := sdk.Coin{
		Denom:  params.DataPrice.Denom,
		Amount: feeAmount,
	}
	return fee, nil
}
