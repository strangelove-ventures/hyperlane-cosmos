package keeper_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
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

func (suite *KeeperTestSuite) mockSetGasOverhead(signer string, igpId uint32, remoteDomain uint32, gasOverhead math.Int) (*types.MsgSetDestinationGasOverheadResponse, error) {
	msg := &types.MsgSetDestinationGasOverhead{
		Sender:            signer,
		DestinationDomain: remoteDomain,
		GasOverhead:       gasOverhead,
		IgpId:             igpId,
	}
	return suite.msgServer.SetDestinationGasOverhead(suite.ctx, msg)
}

func (suite *KeeperTestSuite) mockSetGasPrices(signer string, igpId uint32, remoteDomain uint32, gasPrice math.Int, exchRate math.Int) (*types.MsgSetRemoteGasDataResponse, error) {
	msg := &types.MsgSetRemoteGasData{
		Sender:            signer,
		RemoteDomain:      remoteDomain,
		GasPrice:          gasPrice,
		IgpId:             igpId,
		TokenExchangeRate: exchRate,
	}
	return suite.msgServer.SetRemoteGasData(suite.ctx, msg)
}

func (suite *KeeperTestSuite) mockPayForGas(igpId uint32, signer string, messageId string, remoteDomain uint32, gasAmount math.Int, maxPayment sdk.Coin) (*types.MsgPayForGasResponse, error) {
	msg := &types.MsgPayForGas{
		Sender:            signer,
		MessageId:         messageId,
		DestinationDomain: remoteDomain,
		GasAmount:         gasAmount,
		MaximumPayment:    maxPayment,
		IgpId:             igpId,
	}
	return suite.msgServer.PayForGas(suite.ctx, msg)
}

func (suite *KeeperTestSuite) TestExpectedGasPayments() {
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

func (suite *KeeperTestSuite) TestSetGasPrices() {
	var igpId uint32
	var err error
	var gasPriceMsgSigner string
	var resp *types.MsgSetRemoteGasDataResponse
	oracleAddr := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"
	igpCreator := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	igpBeneficiary := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
	remoteDomain := uint32(1)
	gasPrice := math.NewInt(10000)
	exchangeRate := math.NewInt(10000)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success - oracle sets gas prices",
			func() {
				gasPriceMsgSigner = oracleAddr
				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, remoteDomain)
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"fail - igp owner sets gas overhead",
			func() {
				gasPriceMsgSigner = igpBeneficiary
				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, remoteDomain)
				suite.Require().NoError(err)
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())
			tc.malleate()

			resp, err = suite.mockSetGasPrices(gasPriceMsgSigner, igpId, remoteDomain, gasPrice, exchangeRate)
			events := suite.ctx.EventManager().Events()

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypeGasDataSet,
					sdk.NewAttribute(types.AttributeRemoteDomain, strconv.FormatUint(uint64(remoteDomain), 10)),
					sdk.NewAttribute(types.AttributeTokenExchangeRate, exchangeRate.String()),
					sdk.NewAttribute(types.AttributeGasPrice, gasPrice.String()),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(types.AttributeKeySender, gasPriceMsgSigner),
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

func (suite *KeeperTestSuite) TestSetGasOverhead() {
	var igpId uint32
	var err error
	var gasOhMsgSigner string
	var resp *types.MsgSetDestinationGasOverheadResponse
	oracleAddr := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"
	igpCreator := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	igpBeneficiary := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
	remoteDomain := uint32(1)
	overhead := math.NewInt(10000)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success - igp owner sets gas overhead",
			func() {
				gasOhMsgSigner = igpCreator
				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, remoteDomain)
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"success - oracle sets gas overhead",
			func() {
				gasOhMsgSigner = oracleAddr
				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, remoteDomain)
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"fail - unauthorized address sets gas overhead",
			func() {
				gasOhMsgSigner = igpBeneficiary
				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, remoteDomain)
				suite.Require().NoError(err)
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())
			tc.malleate()

			resp, err = suite.mockSetGasOverhead(gasOhMsgSigner, igpId, remoteDomain, overhead)
			events := suite.ctx.EventManager().Events()

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypeSetGasOverhead,
					sdk.NewAttribute(types.AttributeDestination, strconv.FormatUint(uint64(remoteDomain), 10)),
					sdk.NewAttribute(types.AttributeOverheadAmount, overhead.String()),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(types.AttributeKeySender, gasOhMsgSigner),
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

func (suite *KeeperTestSuite) TestPayForGas() {
	var igpId uint32
	var err error
	var oracleAddr, igpCreator, igpBeneficiary, payer string
	var nativeTokensOwedResp *types.QuoteGasPaymentResponse
	var paymentResp *types.MsgPayForGasResponse
	var exchangeRate math.Int
	var gasPrice math.Int
	var maxPayment sdk.Coin

	testGasAmount := math.NewInt(300000)
	testMessageId := "6ae9a99190641b9ed0c07143340612dde0e9cb7deaa5fe07597858ae9ba5fd7f"
	testDestinationDomain := uint32(1)
	bi := big.NewInt(0)
	var quoteExpected math.Int

	// Most test cases based on hyperlane monorepo test cases found at solidity/test/igps/InterchainGasPaymaster.t.sol
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			// Force an insufficient payment by setting the max allowed payment to one less than the quoteGasPayment (expected amount owed)
			"failure - insufficient payment",
			func() {
				exchangeRate = math.NewInt(1e10)
				gasPrice = math.NewInt(1)
				maxPayment = sdk.NewCoin("stake", math.NewIntFromBigInt(bi))
				quoteExpected = testGasAmount // The exchange rate equals the scale factor (so they divide out to "1", and the price is 1)

				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, testDestinationDomain)
				suite.Require().NoError(err)
				_, err = suite.mockSetGasPrices(oracleAddr, igpId, testDestinationDomain, gasPrice, exchangeRate)
				suite.Require().NoError(err)

				// Get the expected payment amount and denomination
				nativeTokensOwedResp, err = suite.queryClient.QuoteGasPayment(suite.ctx, &types.QuoteGasPaymentRequest{IgpId: igpId, DestinationDomain: testDestinationDomain, GasAmount: testGasAmount})
				suite.Require().NoError(err)
				maxPayment = sdk.NewCoin("stake", math.NewInt(nativeTokensOwedResp.Amount.Int64()-1))
			},
			false,
		},
		{
			// Same exact test case as above, except that the max payment equals the owed payment, so we expect it to pass now
			"success - max payment is expected payment",
			func() {
				exchangeRate = math.NewInt(1e10)
				gasPrice = math.NewInt(1)
				maxPayment = sdk.NewCoin("stake", math.NewIntFromBigInt(bi))
				quoteExpected = testGasAmount // The exchange rate equals the scale factor (so they divide out to "1", and the price is 1)

				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, testDestinationDomain)
				suite.Require().NoError(err)
				_, err = suite.mockSetGasPrices(oracleAddr, igpId, testDestinationDomain, gasPrice, exchangeRate)
				suite.Require().NoError(err)

				// Get the expected payment amount and denomination
				nativeTokensOwedResp, err = suite.queryClient.QuoteGasPayment(suite.ctx, &types.QuoteGasPaymentRequest{IgpId: igpId, DestinationDomain: testDestinationDomain, GasAmount: testGasAmount})
				suite.Require().NoError(err)
				maxPayment = sdk.NewCoin("stake", math.NewInt(nativeTokensOwedResp.Amount.Int64()))
			},
			true,
		},
		{
			// This test case corresponds to the hyperlane monorepo testQuoteGasPaymentRemoteVeryExpensive.
			// 300,000 destination gas
			// 1500 gwei = 1500000000000 wei
			// 300,000 * 1500000000000 = 450000000000000000 (0.45 remote eth)
			// Using the 5000 token exchange rate, meaning the remote native token
			// is 5000x more valuable than the local token:
			// 450000000000000000 * 5000 = 2250000000000000000000 (2250 local eth)

			"success - expensive remote",
			func() {
				// From the hyperlane monorepo, we know what the expected value of the quoteGasPayment is for this test case.
				bi, biSet := bi.SetString("2250000000000000000000", 10)
				suite.Require().True(biSet)
				quoteExpected = math.NewIntFromBigInt(bi)

				exchangeRate = math.NewInt(5000 * 1e10)
				gasPrice = math.NewInt(1500 * 1e9)
				maxPayment = sdk.NewCoin("stake", math.NewIntFromBigInt(bi))

				igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
				suite.Require().NoError(err)
				_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, testDestinationDomain)
				suite.Require().NoError(err)
				_, err = suite.mockSetGasPrices(oracleAddr, igpId, testDestinationDomain, gasPrice, exchangeRate)
				suite.Require().NoError(err)
			},
			true,
		},
	}

	// Create some test addresses initialized with no funds
	testAddrs := simtestutil.AddTestAddrsIncremental(suite.bankKeeper, suite.stakingKeeper, suite.ctx, 3, math.NewInt(0))
	oracleAddr = testAddrs[0].String()
	igpCreator = testAddrs[1].String()
	igpBeneficiary = testAddrs[2].String()

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())
			tc.malleate()

			// Fund the payer with more than enough to pay for the expected gas owed
			payerMinted := quoteExpected.Add(math.NewInt(1000000))
			testAddrPayer := simtestutil.AddTestAddrsIncremental(suite.bankKeeper, suite.stakingKeeper, suite.ctx, 1, payerMinted)
			payer = testAddrPayer[0].String()

			// Sanity check that the balance is what we just funded the account with
			payerCoin := suite.bankKeeper.GetBalance(suite.ctx, testAddrPayer[0], "stake")
			suite.Require().Equal(payerMinted, payerCoin.Amount)

			// Get the expected payment amount and denomination
			nativeTokensOwedResp, err = suite.queryClient.QuoteGasPayment(suite.ctx, &types.QuoteGasPaymentRequest{IgpId: igpId, DestinationDomain: testDestinationDomain, GasAmount: testGasAmount})
			suite.Require().NoError(err)
			suite.Require().Equal(quoteExpected, nativeTokensOwedResp.Amount)

			coinExpected := nativeTokensOwedResp.Amount.String()
			paymentResp, err = suite.mockPayForGas(igpId, payer, testMessageId, testDestinationDomain, testGasAmount, maxPayment)

			events := suite.ctx.EventManager().Events()

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypePayForGas,
					sdk.NewAttribute(types.AttributeMessageId, testMessageId),
					sdk.NewAttribute(types.AttributeGasAmount, testGasAmount.String()),
					sdk.NewAttribute(types.AttributeKeySender, payer),
					sdk.NewAttribute(types.AttributeBeneficiary, igpBeneficiary),
					sdk.NewAttribute(types.AttributePayment, coinExpected),
					sdk.NewAttribute(types.AttributeIgpId, strconv.FormatUint(uint64(igpId), 10)),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(types.AttributeKeySender, payer),
				),
			}

			if tc.expPass {
				for _, evt := range expectedEvents {
					fmt.Printf("TestPayForGas: found event %s", evt.Type)
					suite.Require().Contains(events, evt)
				}

				suite.Require().NoError(err)
				suite.Require().NotNil(paymentResp)

				// Check that the payer's balance was reduced by the gas paid
				payerCoinAfterGasPaid := suite.bankKeeper.GetBalance(suite.ctx, testAddrPayer[0], "stake")
				expectedFundsRemaining := payerMinted.Sub(quoteExpected)
				suite.Require().Equal(expectedFundsRemaining.Int64(), payerCoinAfterGasPaid.Amount.Int64())
			} else {
				suite.Require().Nil(paymentResp)
				suite.Require().Error(err)
			}
		})
	}
}

var ismMap map[uint32]ismtypes.AbstractIsm

func (suite *KeeperTestSuite) SetupIsm() {
	ismMap = map[uint32]ismtypes.AbstractIsm{}

	for _, originIsm := range defaultIsms {
		ism := ismtypes.MustUnpackAbstractIsm(originIsm.AbstractIsm)
		ismMap[originIsm.Origin] = ism
	}
}

func (suite *KeeperTestSuite) VerifyMessage(message []byte, metadata []byte) (bool, error) {
	// Verify Merkle Proof & Validator signature (for LegacyMultiSig). TODO: Ask Steve why this isn't the case for other MultiSigs.
	return ismMap[common.Origin(message)].Verify(metadata, message)
}

func (suite *KeeperTestSuite) TestDispatchPayProcess() {
	var igpId uint32
	var err error
	var oracleAddr, igpCreator, igpBeneficiary, payer string
	var nativeTokensOwedResp *types.QuoteGasPaymentResponse
	var exchangeRate math.Int
	var gasPrice math.Int
	var maxPayment sdk.Coin
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	idx := r1.Intn((len(messages)))
	testGasAmount := math.NewInt(300000)
	testMessageId := common.Id([]byte(messages[idx]))

	testDestinationDomain := uint32(1)
	bi := big.NewInt(0)
	var quoteExpected math.Int

	// Create some test addresses initialized with no funds
	testAddrs := simtestutil.AddTestAddrsIncremental(suite.bankKeeper, suite.stakingKeeper, suite.ctx, 3, math.NewInt(0))
	oracleAddr = testAddrs[0].String()
	igpCreator = testAddrs[1].String()
	igpBeneficiary = testAddrs[2].String()

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success - no setup",
			func() {
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())

			// From the hyperlane monorepo, we know what the expected value of the quoteGasPayment is for this test case.
			bi, biSet := bi.SetString("2250000000000000000000", 10)
			suite.Require().True(biSet)
			quoteExpected = math.NewIntFromBigInt(bi)

			exchangeRate = math.NewInt(5000 * 1e10)
			gasPrice = math.NewInt(1500 * 1e9)
			maxPayment = sdk.NewCoin("stake", math.NewIntFromBigInt(bi))

			igpId, err = suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
			suite.Require().NoError(err)
			_, err = suite.mockCreateOracle(igpCreator, oracleAddr, igpId, testDestinationDomain)
			suite.Require().NoError(err)
			_, err = suite.mockSetGasPrices(oracleAddr, igpId, testDestinationDomain, gasPrice, exchangeRate)
			suite.Require().NoError(err)

			// Fund the payer with more than enough to pay for the expected gas owed
			payerMinted := quoteExpected.Add(math.NewInt(1000000))
			testAddrPayer := simtestutil.AddTestAddrsIncremental(suite.bankKeeper, suite.stakingKeeper, suite.ctx, 1, payerMinted)
			payer = testAddrPayer[0].String()

			// Get the expected payment amount and denomination
			nativeTokensOwedResp, err = suite.queryClient.QuoteGasPayment(suite.ctx, &types.QuoteGasPaymentRequest{IgpId: igpId, DestinationDomain: testDestinationDomain, GasAmount: testGasAmount})
			suite.Require().NoError(err)
			suite.Require().Equal(quoteExpected, nativeTokensOwedResp.Amount)

			coinExpected := nativeTokensOwedResp.Amount.String()
			_, err = suite.mockPayForGas(igpId, payer, string(testMessageId), testDestinationDomain, testGasAmount, maxPayment)

			suite.Require().Contains(suite.ctx.EventManager().Events(), sdk.NewEvent(
				types.EventTypePayForGas,
				sdk.NewAttribute(types.AttributeMessageId, string(testMessageId)),
				sdk.NewAttribute(types.AttributeGasAmount, testGasAmount.String()),
				sdk.NewAttribute(types.AttributeKeySender, payer),
				sdk.NewAttribute(types.AttributeBeneficiary, igpBeneficiary),
				sdk.NewAttribute(types.AttributePayment, coinExpected),
				sdk.NewAttribute(types.AttributeIgpId, strconv.FormatUint(uint64(igpId), 10)),
			))
		})
	}
}
