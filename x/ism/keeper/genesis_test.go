package keeper_test

import (
	"encoding/hex"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	msg := types.NewMsgSetDefaultIsm(signer, defaultIsms)

	_, err := suite.msgServer.SetDefaultIsm(suite.ctx, msg)
	suite.Require().NoError(err)

	gs := suite.keeper.ExportGenesis(suite.ctx)
	for i := 0; i < len(defaultIsms); i++ {
		expectedIsm := types.MustUnpackAbstractIsm(defaultIsms[i].AbstractIsm)
		actualIsm := types.MustUnpackAbstractIsm(gs.DefaultIsm[i].AbstractIsm)
		suite.Require().Equal(defaultIsms[i].Origin, gs.DefaultIsm[i].Origin)
		suite.Require().Equal(expectedIsm, actualIsm)
	}

	suite.SetupTest()

	message, err := hex.DecodeString(messages[0])
	suite.Require().NoError(err)
	metadata, err := hex.DecodeString(metadatas[0])
	suite.Require().NoError(err)
	pass, err := suite.keeper.Verify(metadata, message)
	suite.Require().Error(err)
	suite.Require().False(pass)

	err = suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	pass, err = suite.keeper.Verify(metadata, message)
	suite.Require().NoError(err)
	suite.Require().True(pass)
}
