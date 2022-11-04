package wasmbinding_test

import (
	"testing"
	"time"

	"github.com/CosmosContracts/juno/v11/app"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CreateTestInput(t *testing.T) (*app.App, sdk.Context) {
	junoApp := app.Setup(t, false, 1)
	ctx := junoApp.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "kujira-1", Time: time.Now().UTC()})
	return junoApp, ctx
}

// we need to make this deterministic (same every test run), as content might affect gas costs
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func RandomAccountAddress() sdk.AccAddress {
	_, _, addr := keyPubAddr()
	return addr
}

func RandomBech32AccountAddress() string {
	return RandomAccountAddress().String()
}
