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
	if len(dataSet.Classes) > 0 {
		k.addVip(ctx, buyer)
	}
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
	cost, err := k.calculateCost(ctx, class, buyer)
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
	if len(uploaderSet.UploaderSet) == 0 {
		return fmt.Errorf("uploader set is empty, class: %s", class)
	}
	dataSet, _ := k.getDataByKey(ctx, types.DataSetKey, class)
	totalDataAmount := uint64(len(dataSet.Urls))
	for _, uploader := range uploaderSet.UploaderSet {
		dataSet, _ := k.getDataByKey(ctx, uploader, class)
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

func (k Keeper) calculateCost(ctx sdk.Context, class, buyer string) (sdk.Coin, error) {
	params := k.GetParams(ctx)
	dataSet, _ := k.getDataByKey(ctx, types.DataSetKey, class)
	if len(dataSet.Urls) == 0 {
		return sdk.Coin{}, fmt.Errorf("dataset is empty, class: %s", class)
	}
	dataAmount := int64(len(dataSet.Urls))
	feeAmount := params.DataPrice.Amount.MulRaw(dataAmount)
	vipInfo, _ := k.getVipInfo(ctx)
	if goutil.Contains(vipInfo.VipList, buyer) {
		feeAmount = math.NewInt(int64(float64(feeAmount.Uint64()) * vipInfo.Discount))
	}
	fee := sdk.Coin{
		Denom:  params.DataPrice.Denom,
		Amount: feeAmount,
	}
	return fee, nil
}

func (k Keeper) MintTo(ctx sdk.Context, amount sdk.Coin, to string) error {
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}

	addr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return err
	}

	if k.bankKeeper.BlockedAddr(addr) {
		return fmt.Errorf("failed to mint to blocked address: %s", addr)
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(amount))
}

func (k Keeper) UpdateVipDiscount(ctx sdk.Context, discount float64) error {
	store := k.GetVipInfoPrefixStore(ctx)

	vipInfo, _ := k.getVipInfo(ctx)
	vipInfo.Discount = discount
	bz, err := proto.Marshal(&vipInfo)
	if err != nil {
		return err
	}
	store.Set([]byte(types.VipInfoKey), bz)
	return nil
}

func (k Keeper) addVip(ctx sdk.Context, addr string) error {
	store := k.GetVipInfoPrefixStore(ctx)

	vipInfo, _ := k.getVipInfo(ctx)
	if goutil.Contains(vipInfo.VipList, addr) {
		return nil
	}
	vipInfo.VipList = append(vipInfo.VipList, addr)
	bz, err := proto.Marshal(&vipInfo)
	if err != nil {
		return err
	}
	store.Set([]byte(types.VipInfoKey), bz)
	return nil
}

func (k Keeper) getVipInfo(ctx sdk.Context) (types.VipInfo, error) {
	store := k.GetVipInfoPrefixStore(ctx)

	bz := store.Get([]byte(types.VipInfoKey))
	vipInfo := types.VipInfo{}
	err := proto.Unmarshal(bz, &vipInfo)
	if err != nil {
		return types.VipInfo{}, err
	}
	return vipInfo, nil
}
