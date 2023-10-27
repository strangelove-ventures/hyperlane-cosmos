package keeper_test

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

func (suite *KeeperTestSuite) TestQueryOriginsDefaultIsm() {
	var req *types.QueryOriginsDefaultIsmRequest
	var index uint32

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success with index 0",
			func() {
				index = 0
				req = &types.QueryOriginsDefaultIsmRequest{Origin: defaultIsms[index].Origin}
			},
			true,
		},
		{
			"success with index 1",
			func() {
				index = 1
				req = &types.QueryOriginsDefaultIsmRequest{Origin: defaultIsms[index].Origin}
			},
			true,
		},
		{
			"success with index 2",
			func() {
				index = 2
				req = &types.QueryOriginsDefaultIsmRequest{Origin: defaultIsms[index].Origin}
			},
			true,
		},
		{
			"fails with empty request",
			func() {
				req = &types.QueryOriginsDefaultIsmRequest{}
			},
			false,
		},
		// TODO: add test case with invalid origin / no default ISM set
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
			msg := types.NewMsgSetDefaultIsm(signer, defaultIsms)

			_, err := suite.keeper.SetDefaultIsm(suite.ctx, msg)
			suite.Require().NoError(err)

			tc.malleate()

			res, err := suite.queryClient.OriginsDefaultIsm(suite.ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(res.DefaultIsm)
				suite.Require().Equal(defaultIsms[index].AbstractIsm.String(), res.DefaultIsm.String())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryAllDefaultIsms() {
	var req *types.QueryAllDefaultIsmsRequest

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				req = &types.QueryAllDefaultIsmsRequest{}
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
			msg := types.NewMsgSetDefaultIsm(signer, defaultIsms)

			_, err := suite.keeper.SetDefaultIsm(suite.ctx, msg)
			suite.Require().NoError(err)

			tc.malleate()

			res, err := suite.queryClient.AllDefaultIsms(suite.ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(res.DefaultIsms)
				suite.Require().Equal(defaultIsms[0].String(), res.DefaultIsms[0].String())
				suite.Require().Equal(defaultIsms[1].String(), res.DefaultIsms[1].String())
				suite.Require().Equal(defaultIsms[2].String(), res.DefaultIsms[2].String())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}


func (suite *KeeperTestSuite) TestQueryCustomIsm() {
	var req *types.QueryCustomIsmRequest
	var index uint32

	signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success with index 1",
			func() {
				index = 1
				req = &types.QueryCustomIsmRequest{IsmId: index}
				msg := types.NewMsgCreateIsm(signer, defaultIsms[index-1].AbstractIsm)
				_, err := suite.keeper.CreateIsm(suite.ctx, msg)
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"success with index 1",
			func() {
				index = 2
				req = &types.QueryCustomIsmRequest{IsmId: index}
				msg := types.NewMsgCreateIsm(signer, defaultIsms[index-1].AbstractIsm)
				_, err := suite.keeper.CreateIsm(suite.ctx, msg)
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"success with index 3",
			func() {
				index = 3
				req = &types.QueryCustomIsmRequest{IsmId: index}
				msg := types.NewMsgCreateIsm(signer, defaultIsms[index-1].AbstractIsm)
				_, err := suite.keeper.CreateIsm(suite.ctx, msg)
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"fails with empty request",
			func() {
				req = &types.QueryCustomIsmRequest{}
			},
			false,
		},
		{
			"fails with index 0",
			func() {
				index = 0
				req = &types.QueryCustomIsmRequest{IsmId: index}
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			
			tc.malleate()
			res, err := suite.queryClient.CustomIsm(suite.ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(res.CustomIsm)
				suite.Require().Equal(defaultIsms[index-1].AbstractIsm.String(), res.CustomIsm.String())
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryAllCustomIsms() {
	var req *types.QueryAllCustomIsmsRequest

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				req = &types.QueryAllCustomIsmsRequest{}
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
			
			msg := types.NewMsgCreateIsm(signer, defaultIsms[0].AbstractIsm)
			resp1, err := suite.keeper.CreateIsm(suite.ctx, msg)
			suite.Require().NoError(err)

			msg = types.NewMsgCreateIsm(signer, defaultIsms[1].AbstractIsm)
			resp2, err := suite.keeper.CreateIsm(suite.ctx, msg)
			suite.Require().NoError(err)

			msg = types.NewMsgCreateIsm(signer, defaultIsms[2].AbstractIsm)
			resp3, err := suite.keeper.CreateIsm(suite.ctx, msg)
			suite.Require().NoError(err)

			tc.malleate()

			res, err := suite.queryClient.AllCustomIsms(suite.ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(resp1.IsmId, uint32(1))
				suite.Require().Equal(resp2.IsmId, uint32(2))
				suite.Require().Equal(resp3.IsmId, uint32(3))
				suite.Require().NotEmpty(res.CustomIsms)
				suite.Require().Equal(defaultIsms[0].AbstractIsm.String(), res.CustomIsms[0].AbstractIsm.String())
				suite.Require().Equal(defaultIsms[1].AbstractIsm.String(), res.CustomIsms[1].AbstractIsm.String())
				suite.Require().Equal(defaultIsms[2].AbstractIsm.String(), res.CustomIsms[2].AbstractIsm.String())
				suite.Require().Equal(uint32(1), res.CustomIsms[0].Index)
				suite.Require().Equal(uint32(2), res.CustomIsms[1].Index)
				suite.Require().Equal(uint32(3), res.CustomIsms[2].Index)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

