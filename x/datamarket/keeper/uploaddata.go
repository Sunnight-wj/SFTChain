package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/gookit/goutil"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

func (k Keeper) UploadData(ctx sdk.Context, class, sender, url string) error {
	if err := k.addData(ctx, class, sender, url); err != nil {
		return err
	}
	return nil
}

func (k Keeper) addData(ctx sdk.Context, uploader, class, url string) error {
	err := k.addDataByKey(ctx, types.DataSetKey, class, url)
	if err != nil {
		return err
	}
	err = k.addDataByKey(ctx, uploader, class, url)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) addDataByKey(ctx sdk.Context, key, class, url string) error {
	dataSet, _ := k.getDataByKey(ctx, key, class)
	if goutil.Contains(dataSet.Urls, url) {
		return fmt.Errorf("data: %s already exist in dataset: %s", url, class)
	}
	dataSet.Urls = append(dataSet.Urls, url)
	bz, err := proto.Marshal(&dataSet)
	if err != nil {
		return err
	}
	store := k.GetDataSetPrefixStore(ctx, class)
	store.Set([]byte(key), bz)
	return nil
}

func (k Keeper) getDataByKey(ctx sdk.Context, key, class string) (types.DataSet, error) {
	bz := k.GetDataSetPrefixStore(ctx, class).Get([]byte(key))

	dataSet := types.DataSet{}
	err := proto.Unmarshal(bz, &dataSet)
	if err != nil {
		return types.DataSet{}, err
	}
	return dataSet, nil
}
