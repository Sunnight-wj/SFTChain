package keeper

import (
	"fmt"
	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper

		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	authority string,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		authority:     authority,
	}
}

func (k Keeper) getAuthority() string {
	return k.authority
}

type DataMarket string

func (d DataMarket) Name() string {
	return "datamarket"
}

func (d DataMarket) String() string {
	return "datamarket"
}

func Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateModuleAccount(ctx sdk.Context) {
	moduleAcc := authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter, authtypes.Burner)
	k.accountKeeper.SetModuleAccount(ctx, moduleAcc)
}

func (k Keeper) GetDataSetPrefixStore(ctx sdk.Context, class string) sdk.KVStore {
	key := DataMarket("datamarket")
	store := ctx.KVStore(key)
	return prefix.NewStore(store, types.GetDataSetPrefix(class))
}

func (k Keeper) GetBuyerPrefixStore(ctx sdk.Context, buyer string) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, types.GetBuyerPrefix(buyer))
}
