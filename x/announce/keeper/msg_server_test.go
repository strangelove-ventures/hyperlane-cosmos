package keeper_test

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func (suite *KeeperTestSuite) TestAnnouncement() {
	var msg *types.MsgAnnouncement
	txSigner := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	valPrivKey := "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61" // Testing only, do NOT use
	validatorPrivateKey, err := crypto.HexToECDSA(valPrivKey)
	suite.Require().NoError(err)
	valAddr := crypto.PubkeyToAddress(validatorPrivateKey.PublicKey)
	var validator *counterchain.CounterChain
	storageLocation := "file:///tmp//signatures-simd1"

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				validator = counterchain.CreateEmperorValidator(suite.T(), suite.mailboxKeeper.GetDomain(suite.ctx), counterchain.LEGACY_MULTISIG, valPrivKey)
				digest, err := types.GetAnnouncementDigest(suite.mailboxKeeper.GetDomain(suite.ctx), suite.mailboxKeeper.GetMailboxAddress(), storageLocation)
				suite.Require().NoError(err)
				valSignature := validator.Sign(digest)

				msg = &types.MsgAnnouncement{ // Valid announcement with valid signature
					Sender:          txSigner,
					Validator:       valAddr.Bytes(),
					StorageLocation: storageLocation,
					Signature:       valSignature,
				}
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest(suite.T())
			tc.malleate()

			resp, err := suite.msgServer.Announcement(suite.ctx, msg)
			suite.Require().NoError(err)
			events := suite.ctx.EventManager().Events()
			valAddrHex := hex.EncodeToString(msg.Validator)

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypeAnnounce,
					sdk.NewAttribute(types.AttributeStorageLocation, storageLocation),
					sdk.NewAttribute(types.AttributeValidatorAddress, valAddrHex),
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
