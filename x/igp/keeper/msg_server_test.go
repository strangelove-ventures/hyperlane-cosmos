package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

func (suite *KeeperTestSuite) TestCreateIgp() {
	var (
		msg *types.MsgCreateIgp
		id  uint32
	)

	sender := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	recipient := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				msg = &types.MsgCreateIgp{
					Sender:                 sender,
					Beneficiary:            recipient,
					TokenExchangeRateScale: math.ZeroInt(),
				}
				id = 0
			},
			true,
		},
		{
			"fail",
			func() {
				msg = &types.MsgCreateIgp{
					Sender:                 sender,
					Beneficiary:            recipient,
					TokenExchangeRateScale: math.NewInt(-1), // Scale is not allowed (Must be 0 to indicate the scale wasn't provided, or >= 1)
				}
				id = 1
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())

			tc.malleate()

			resp, err := suite.msgServer.CreateIgp(suite.ctx, msg)
			events := suite.ctx.EventManager().Events()

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypeCreateIgp,
					sdk.NewAttribute(types.AttributeIgpId, "0"),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
				),
			}

			if tc.expPass {
				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}

				suite.Require().NoError(err)
				suite.Require().NotNil(resp)
				suite.Require().Equal(id, resp.IgpId)
			} else {
				suite.Require().Nil(resp)
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) mockCreateIgp(sender string, recipient string, scale math.Int) (id uint32, err error) {
	msg := &types.MsgCreateIgp{
		Sender:                 sender,
		Beneficiary:            recipient,
		TokenExchangeRateScale: scale,
	}
	resp, err := suite.msgServer.CreateIgp(suite.ctx, msg)
	if err != nil {
		return 0, err
	}

	return resp.IgpId, err
}

func (suite *KeeperTestSuite) mockCreateOracle(signer string, gasOracleAddr string, igpId uint32, remoteDomain uint32) (*types.MsgSetGasOraclesResponse, error) {
	msg := &types.MsgSetGasOracles{
		Sender: signer,
		Configs: []*types.GasOracleConfig{
			{
				IgpId:        igpId,
				GasOracle:    gasOracleAddr,
				RemoteDomain: remoteDomain,
			},
		},
	}
	return suite.msgServer.SetGasOracles(suite.ctx, msg)
}

func (suite *KeeperTestSuite) TestCreateOracle() {
	var igpId uint32
	var err error
	var signer string
	var oracleCreateOrUpdateEvt sdk.Event

	igpBeneficiary := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
	oracleAddr := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"
	remoteDomain := uint32(1)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success - create new Oracle",
			func() {
				signer = "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
				igpId, err = suite.mockCreateIgp(signer, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				oracleCreateOrUpdateEvt = sdk.NewEvent(
					types.EventTypeCreateOracle,
					sdk.NewAttribute(types.AttributeOracleAddress, oracleAddr),
				)
			},
			true,
		},
		{
			"success - update existing Oracle",
			func() {
				signer = "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
				igpId, err = suite.mockCreateIgp(signer, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err := suite.mockCreateOracle(signer, oracleAddr, igpId, remoteDomain)
				suite.Require().NoError(err)

				// Confirms that the oracle was created and persisted to storage (causing an update event on second call to mockCreateOracle)
				oracleCreateOrUpdateEvt = sdk.NewEvent(
					types.EventTypeUpdateOracle,
					sdk.NewAttribute(types.AttributeOracleAddress, oracleAddr),
				)
			},
			true,
		},
		{
			"failed, address unauthorized to update oracle",
			func() {
				signer = "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
				igpId, err = suite.mockCreateIgp(signer, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				signer = oracleAddr // This will cause an unauthorized response; oracle cannot update an IGP it did not create
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())
			tc.malleate()

			oracleResp, err := suite.mockCreateOracle(signer, oracleAddr, igpId, remoteDomain)

			events := suite.ctx.EventManager().Events()

			// Verify events
			expectedEvents := sdk.Events{
				oracleCreateOrUpdateEvt,
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(types.AttributeKeySender, signer),
				),
			}

			if tc.expPass {
				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}

				suite.Require().NoError(err)
				suite.Require().NotNil(oracleResp)
			} else {
				suite.Require().Nil(oracleResp)
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSetBeneficiary() {
	var msg *types.MsgSetBeneficiary

	igpCreator := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	igpBeneficiary := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
	addrOther := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				msg = &types.MsgSetBeneficiary{ // Case 1: IGP Creator is allowed to update the beneficiary
					Sender:  igpCreator,
					Address: addrOther,
				}
			},
			true,
		},
		{
			"fail",
			func() {
				msg = &types.MsgSetBeneficiary{ // Case 2: Only the IGP owner can update the beneficiary
					Sender:  addrOther,
					Address: igpBeneficiary,
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())
			tc.malleate()

			igpId, err := suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
			suite.Require().NoError(err)

			// Update the MsgSetBeneficiary with the IGP ID that was just created
			msg.IgpId = igpId
			resp, err := suite.msgServer.SetBeneficiary(suite.ctx, msg)
			events := suite.ctx.EventManager().Events()

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypeSetBeneficiary,
					sdk.NewAttribute(types.AttributeIgpId, "0"),
					sdk.NewAttribute(types.AttributeBeneficiary, msg.Address),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
				),
			}

			if tc.expPass {
				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}

				suite.Require().NoError(err)
				suite.Require().NotNil(resp)
			} else {
				suite.Require().Nil(resp)
				suite.Require().Error(err)
			}
		})
	}
}
