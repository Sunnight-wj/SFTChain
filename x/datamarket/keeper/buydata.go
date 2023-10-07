package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/gookit/goutil"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

func (k Keeper) BuyData(ctx sdk.Context, buyer, class string) error {
	// 付费
	err := k.chargeForBuyData(ctx, buyer, class)
	if err != nil {
		return err
	}
	// 写数据库
	err = k.addDataSetForBuyer(ctx, buyer, class)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) addDataSetForBuyer(ctx sdk.Context, buyer, class string) error {
	dataSet, _ := k.getDataSetByBuyer(ctx, buyer)
	if goutil.Contains(dataSet.Classes, class) {
		return fmt.Errorf("buyer already buy the class: %s", class)
	}
	dataSet.Classes = append(dataSet.Classes, class)
	bz, err := proto.Marshal(&dataSet)
	if err != nil {
		return err
	}
	store := k.GetBuyerPrefixStore(ctx, buyer)
	store.Set([]byte(types.BuyerKey), bz)
	return nil
}

func (k Keeper) getDataSetByBuyer(ctx sdk.Context, buyer string) (types.Classes, error) {
	store := k.GetBuyerPrefixStore(ctx, buyer)

	bz := store.Get([]byte(types.BuyerKey))
	classes := types.Classes{}
	err := proto.Unmarshal(bz, &classes)
	if err != nil {
		return types.Classes{}, err
	}
	return classes, nil
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
	k.distributeCost(ctx, class, cost)
	return nil
}

func (k Keeper) distributeCost(ctx sdk.Context, class string, cost sdk.Coin) error {
	feePercentage := k.GetParams(ctx).FeePercentage
	costAmount := cost.Amount.Uint64()
	feeAmount := uint64(float64(costAmount) * feePercentage)
	rewardAmount := costAmount - feeAmount
	uploaderSet, err := k.getUploaderSet(ctx, class)
	if err != nil {
		return err
	}
	dataSet, _ := k.getDataByKey(ctx, types.DataSetKey, class)
	totalDataAmount := uint64(len(dataSet.Urls))
	for _, uploader := range uploaderSet.UploaderSet {
		dataSet, _ := k.getDataByKey(ctx, class, uploader)
		dataAmount := uint64(len(dataSet.Urls))
		dataCost := dataAmount * rewardAmount / totalDataAmount
		uploaderReward := sdk.Coin{
			Denom:  cost.Denom,
			Amount: math.NewInt(int64(dataCost)),
		}
		uploaderAddr, _ := sdk.AccAddressFromBech32(uploader)
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, uploaderAddr, []sdk.Coin{uploaderReward})
	}
	return nil
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
