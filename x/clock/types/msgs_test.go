package types

import (
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v17/x/drip/types"
)

type MsgsTestSuite struct {
	suite.Suite
	govModule string
}

func TestMsgsTestSuite(t *testing.T) {
	suite.Run(t, new(MsgsTestSuite))
}

func (suite *MsgsTestSuite) SetupTest() {
	suite.govModule = "juno10d07y265gmmuvt4z0w9aw880jnsr700jvss730"
}

func (suite *MsgsTestSuite) TestMsgUpdateParams() {
	p := MsgUpdateParams{
		Authority: suite.govModule,
		Params: Params{
			ContractAddresses: []string{},
		},
	}

	acc, _ := sdk.AccAddressFromBech32(p.Authority)

	msg := NewMsgUpdateParams(acc, []string(nil))

	suite.Require().Equal(types.RouterKey, msg.Route())
	suite.Require().Equal(TypeMsgUpdateParams, msg.Type())
	suite.Require().NotNil(msg.GetSigners())
}
