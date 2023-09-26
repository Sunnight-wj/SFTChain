package keeper

import (
	"fmt"
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

func (k Keeper) addData(ctx sdk.Context, class, sender, url string) error {
	url, err := k.getData(ctx, class, sender)
	if err == nil {
		return fmt.Errorf("data already set for class: %s, sender: %s", class, sender)
	}
	store := k.GetClassPrefixStore(ctx, class)
	store.Set([]byte(sender), []byte(url))
	return nil
}

func (k Keeper) getData(ctx sdk.Context, class, sender string) (string, error) {
	store := k.GetClassPrefixStore(ctx, class)
	if url := store.Get([]byte(sender)); url != nil {
		return string(url), nil
	}
	return "", fmt.Errorf("data not found for class: %s, sender: %s", class, sender)
}
