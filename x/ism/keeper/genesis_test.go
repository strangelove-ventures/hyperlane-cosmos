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

	customIsmMsg1 := types.NewMsgCreateIsm(signer, defaultIsms[0].AbstractIsm)
	_, err = suite.msgServer.CreateIsm(suite.ctx, customIsmMsg1)
	suite.Require().NoError(err)

	customIsmMsg2 := types.NewMsgCreateIsm(signer, defaultIsms[1].AbstractIsm)
	_, err = suite.msgServer.CreateIsm(suite.ctx, customIsmMsg2)
	suite.Require().NoError(err)

	customIsmMsg3 := types.NewMsgCreateIsm(signer, defaultIsms[2].AbstractIsm)
	_, err = suite.msgServer.CreateIsm(suite.ctx, customIsmMsg3)
	suite.Require().NoError(err)

	gs := suite.keeper.ExportGenesis(suite.ctx)
	for i := 0; i < len(defaultIsms); i++ {
		expectedIsm := types.MustUnpackAbstractIsm(defaultIsms[i].AbstractIsm)
		actualDefaultIsm := types.MustUnpackAbstractIsm(gs.DefaultIsm[i].AbstractIsm)
		suite.Require().Equal(defaultIsms[i].Origin, gs.DefaultIsm[i].Origin)
		suite.Require().Equal(expectedIsm, actualDefaultIsm)
		
		actualCustomIsm := types.MustUnpackAbstractIsm(gs.CustomIsm[i].AbstractIsm)
		suite.Require().Equal(uint32(i+1), gs.CustomIsm[i].Index)
		suite.Require().Equal(expectedIsm, actualCustomIsm)
	}

	suite.SetupTest()

	message, err := hex.DecodeString(messages[0])
	suite.Require().NoError(err)
	metadata, err := hex.DecodeString(metadatas[0])
	suite.Require().NoError(err)
	pass, err := suite.keeper.Verify(suite.ctx, metadata, message, 0)
	suite.Require().Error(err)
	suite.Require().False(pass)

	err = suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	pass, err = suite.keeper.Verify(suite.ctx, metadata, message, 0)
	suite.Require().NoError(err)
	suite.Require().True(pass)

	for i := 0; i < len(defaultIsms); i++ {
		expectedIsm := types.MustUnpackAbstractIsm(defaultIsms[i].AbstractIsm)

		defaultIsmReq := types.QueryOriginsDefaultIsmRequest{
			Origin: defaultIsms[i].Origin,
		}
		defaultIsmResp, err := suite.queryClient.OriginsDefaultIsm(suite.ctx, &defaultIsmReq)
		suite.Require().NoError(err)
		actualDefaultIsm := types.MustUnpackAbstractIsm(defaultIsmResp.DefaultIsm)
		suite.Require().Equal(defaultIsms[i].Origin, gs.DefaultIsm[i].Origin)
		suite.Require().Equal(expectedIsm, actualDefaultIsm)
		
		customIsmReq := types.QueryCustomIsmRequest{
			IsmId: uint32(i+1),
		}
		customIsmResp, err := suite.queryClient.CustomIsm(suite.ctx, &customIsmReq)
		suite.Require().NoError(err)
		actualCustomIsm := types.MustUnpackAbstractIsm(customIsmResp.CustomIsm)
		suite.Require().Equal(uint32(i+1), gs.CustomIsm[i].Index)
		suite.Require().Equal(expectedIsm, actualCustomIsm)
	}
}
