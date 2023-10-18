package v18

import (
	"github.com/CosmosContracts/juno/v17/app/keepers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		registerJunoMetadata(ctx, keepers.BankKeeper)
		return mm.RunMigrations(ctx, cfg, fromVM)
	}
}

func registerJunoMetadata(ctx sdk.Context, bankKeeper bankkeeper.Keeper) {
	ujunoMetadata := banktypes.Metadata{
		Description: "The native token of Juno",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "ujuno",
				Exponent: 0,
				Aliases:  nil,
			},
			{
				Denom:    "juno",
				Exponent: 6,
				Aliases:  nil,
			},
		},
		Base:    "ujuno",
		Display: "juno",
	}
	bankKeeper.SetDenomMetaData(ctx, ujunoMetadata)
}
