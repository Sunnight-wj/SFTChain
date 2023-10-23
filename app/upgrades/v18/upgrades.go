package v18

import (
	"fmt"
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
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// Run migrations
		logger.Info(fmt.Sprintf("pre migrate version map: %v", vm))
		versionMap, err := mm.RunMigrations(ctx, cfg, vm)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		registerJunoMetadata(ctx, keepers.BankKeeper)
		return vm, err
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
