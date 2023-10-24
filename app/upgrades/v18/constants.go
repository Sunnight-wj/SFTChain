package v18

import (
	"github.com/CosmosContracts/juno/v17/app/upgrades"
	hellotypes "github.com/CosmosContracts/juno/v17/x/hello/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

// UpgradeName defines the on-chain upgrade name for the upgrade.
const UpgradeName = "v18"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			// new modules
			hellotypes.ModuleName,
		},
	},
}
