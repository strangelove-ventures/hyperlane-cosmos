package keeper_test

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func (suite *KeeperTestSuite) TestDispatch() {
	var (
		msg    *types.MsgDispatch
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				sender := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
				//sender := "cosmos17lklk2z0gevmdx4tvjac7g7t9prqczg6p6nuw9"
				recipient := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
				domain := uint32(12)
				msgBody := "Hello"
				msg = types.NewMsgDispatch(sender, domain, recipient, msgBody)
			},
			true,
		},
		/*{
			"fails with unauthorized signer",
			func() {
				msg = types.NewMsgSetDefaultIsm(signer, defaultIsms)
			},
			false,
		},*/
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			tc.malleate()

			res, err := suite.msgServer.Dispatch(suite.ctx, msg)
			events := suite.ctx.EventManager().Events()

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				// Verify events
				expectedEvents := sdk.Events{
					sdk.NewEvent(
						types.EventTypeDispatch,
						sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
						sdk.NewAttribute(types.AttributeKeyDestinationDomain, strconv.FormatUint(uint64(msg.DestinationDomain), 10)),
						sdk.NewAttribute(types.AttributeKeyRecipientAddress, msg.RecipientAddress),
						sdk.NewAttribute(types.AttributeKeyMessage, msg.MessageBody),
					),
					sdk.NewEvent(
						types.EventTypeDispatchId,
						sdk.NewAttribute(types.AttributeKeyID, res.MessageId),
					),
					sdk.NewEvent(
						sdk.EventTypeMessage,
						sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
					),
				}

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
