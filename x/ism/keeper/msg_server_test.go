package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

func (suite *KeeperTestSuite) TestMsgSetDefaultIsm() {
	var (
		msg    *types.MsgSetDefaultIsm
		signer string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				msg = types.NewMsgSetDefaultIsm(signer, defaultIsms)
			},
			true,
		},
		{
			"fails with unauthorized signer",
			func() {
				signer = authtypes.NewModuleAddress(authtypes.ModuleName).String()
				msg = types.NewMsgSetDefaultIsm(signer, defaultIsms)
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			signer = authtypes.NewModuleAddress(govtypes.ModuleName).String()

			tc.malleate()

			err := msg.ValidateBasic()
			suite.Require().NoError(err)

			res, err := suite.msgServer.SetDefaultIsm(suite.ctx, msg)
			events := suite.ctx.EventManager().Events()

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				// Verify events
				expectedEvents := sdk.Events{}
				for _, originIsm := range defaultIsms {
					ism := types.MustUnpackAbstractIsm(originIsm.AbstractIsm)
					expectedEvents.AppendEvent(ism.Event(originIsm.Origin))
				}
				expectedEvents.AppendEvent(sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				))
				
				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
				suite.Require().Empty(events)
			}
		})
	}
}
