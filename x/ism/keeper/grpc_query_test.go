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

// TODO: add test case for querying all ISMs