package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	_ "github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

func (k Keeper) UploadData(ctx sdk.Context, class, sender, url string) error {
	err := k.addData(ctx, class, sender, url)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) addData(ctx sdk.Context, uploader, class, url string) error {
	store := k.GetDataSetPrefixStore(ctx, class)
	store.Set([]byte(class), )

}

func (k Keeper) getData(ctx sdk.Context, class string) ([]string, error) {

}
