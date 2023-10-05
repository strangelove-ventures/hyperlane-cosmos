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

	validator := counterchain.CreateEmperorValidator(suite.T(), uint32(suite.domain), counterchain.LEGACY_MULTISIG, valPrivKey)
	storageLocation := "file:///tmp//signatures-simd1"
	digest, err := types.GetAnnouncementDigest(suite.domain, suite.mailboxKeeper.GetMailboxAddress(), storageLocation)
	suite.Require().NoError(err)
	valSignature := validator.Sign(digest)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				msg = &types.MsgAnnouncement{ // Case 1: IGP Creator is allowed to update the beneficiary
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

			//_, err := suite.mockAnnounce()
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
			}

			if tc.expPass {
				for _, evt := range expectedEvents {
					suite.Require().Contains(events, evt)
				}

				suite.Require().NoError(err)
				//suite.Require().NotNil(resp)
			} else {
				//suite.Require().Nil(resp)
				suite.Require().Error(err)
			}
		})
	}
}
