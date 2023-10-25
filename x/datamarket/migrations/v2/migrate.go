package v2

import (
	"fmt"
	v2types "github.com/CosmosContracts/juno/v17/x/datamarket/migrations/v2/types"
	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "datamarket"
)

func Migrate(
	_ sdk.Context,
	store sdk.KVStore,
	cdc codec.BinaryCodec,
) error {
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.DataSetPrefix))
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		val := iterator.Value()
		var oldDataSet v2types.DataSet
		cdc.MustUnmarshal(val, &oldDataSet)
		class := string(key)
		newDataSet := types.DataSet{
			Urls:        oldDataSet.Urls,
			Description: fmt.Sprintf("this data set belong to %s", class),
		}
		bz := cdc.MustMarshal(&newDataSet)
		store.Set(key, bz)
	}
	return nil
}
