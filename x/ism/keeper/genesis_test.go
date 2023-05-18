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
		suite.Require().Equal(defaultIsms[i].Origin, gs.DefaultIsm[i].Origin)
		suite.Require().Equal(defaultIsms[i].Ism.Threshold, gs.DefaultIsm[i].Ism.Threshold)
		for j := 0; j < len(defaultIsms[i].Ism.ValidatorPubKeys); j++ {
			suite.Require().Equal(defaultIsms[i].Ism.ValidatorPubKeys[j], gs.DefaultIsm[i].Ism.ValidatorPubKeys[j])
		}
	}

	suite.SetupTest()

	message, err := hex.DecodeString(messages[0])
	suite.Require().NoError(err)
	metadata, err := hex.DecodeString(metadatas[0])
	suite.Require().NoError(err)
	pass := suite.keeper.Verify(metadata, message)
	suite.Require().False(pass)

	err = suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	pass = suite.keeper.Verify(metadata, message)
	suite.Require().True(pass)
}
