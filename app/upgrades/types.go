package upgrades

import (
	"strings"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/CosmosContracts/juno/v17/app/keepers"
)

// BaseAppParamManager defines an interrace that BaseApp is expected to fullfil
// that allows upgrade handlers to modify BaseApp parameters.
type BaseAppParamManager interface {
	GetConsensusParams(ctx sdk.Context) *tmproto.ConsensusParams
	StoreConsensusParams(ctx sdk.Context, cp *tmproto.ConsensusParams)
}

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(
		*module.Manager,
		module.Configurator,
		*keepers.AppKeepers,
	) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades store.StoreUpgrades
}

// Returns "ujunox" if the chain is uni, else returns the standard ujuno token denom.
func GetChainsDenomToken(chainID string) string {
	if strings.HasPrefix(chainID, "uni-") {
		return "ujunox"
	}
	return "ujuno"
}
