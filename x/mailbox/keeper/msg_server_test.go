package keeper_test

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func (suite *KeeperTestSuite) TestDispatch() {
	var (
		msg *types.MsgDispatch
		id  string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success with contract as sender",
			func() {
				id = "0x806d81a5b017cc51297bb545bc037f39ceecefab8041766c9b733900c9a01242"
				sender := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
				recipientBech32 := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
				recipientBytes := sdk.MustAccAddressFromBech32(recipientBech32).Bytes()
				recipientHex := hexutil.Encode(recipientBytes)
				domain := uint32(12)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(sender, domain, recipientHex, msgBody)
			},
			true,
		},
		{
			"success with account as sender",
			func() {
				id = "0x0c4024fb3262d8852b1fc5caa9f73f91b9375a1cbe51f6da52d192d272623fd1"
				sender := "cosmos17lklk2z0gevmdx4tvjac7g7t9prqczg6p6nuw9"
				recipientBech32 := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
				recipientBytes := sdk.MustAccAddressFromBech32(recipientBech32).Bytes()
				recipientHex := hexutil.Encode(recipientBytes)
				domain := uint32(12)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(sender, domain, recipientHex, msgBody)
			},
			true,
		},
		{
			"success with hyperlane explorer data w/ sender/recipient padding",
			func() {
				id = "0x82fbf348d68e34903627d1f57fa3227211e35134e2a613ce6b78ce2d42e17198"
				senderHex := "0x0000000000000000000000002b0860e52244f03e59f12cfe413d6a29bc30b893"
				senderBytes := hexutil.MustDecode(senderHex)
				senderBech32 := sdk.MustBech32ifyAddressBytes("cosmos", senderBytes)
				recipient := "0x00000000000000000000000076a2f655352752af6ce9b03932b9090009dc5d0c"
				domain := uint32(43114)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(senderBech32, domain, recipient, msgBody)
			},
			true,
		},
		{
			"success with hyperlane explorer data w/o sender/recipient padding",
			func() {
				id = "0x82fbf348d68e34903627d1f57fa3227211e35134e2a613ce6b78ce2d42e17198"
				senderHex := "0x2b0860e52244f03e59f12cfe413d6a29bc30b893"
				senderBytes := hexutil.MustDecode(senderHex)
				senderBech32 := sdk.MustBech32ifyAddressBytes("cosmos", senderBytes)
				recipient := "0x76a2f655352752af6ce9b03932b9090009dc5d0c"
				domain := uint32(43114)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(senderBech32, domain, recipient, msgBody)
			},
			true,
		},
		{
			"fails with wrong id match",
			func() {
				id = "0x82fbf348d68e34903627d1f57fa3227211e35134e2a613ce6b78ce2d42e17198"
				senderHex := "0x2b0860e52244f03e59f12cfe413d6a29bc30b893"
				senderBytes := hexutil.MustDecode(senderHex)
				senderBech32 := sdk.MustBech32ifyAddressBytes("cosmos", senderBytes)
				recipient := "0x76a2f655352752af6ce9b03932b9090009dc5d0c"
				domain := uint32(43114)
				msgBody := hexutil.Encode([]byte("Hello")) // Removed !
				msg = types.NewMsgDispatch(senderBech32, domain, recipient, msgBody)
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			tc.malleate()

			res, err := suite.msgServer.Dispatch(suite.ctx, msg)
			events := suite.ctx.EventManager().Events()

			suite.Require().NoError(err)
			suite.Require().NotNil(res)

			fmt.Println("ID: ", res.MessageId)

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
			if tc.expPass {
				suite.Require().Equal(id, res.MessageId)
			} else {
				suite.Require().NotEqual(id, res.MessageId)
			}
		})
	}
}

/*func (suite *KeeperTestSuite) TestProcessMsg() {
	counterChain := helpers.CreateCounterChain(suite.T(), 112)
	// Set default ISM
	signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	var valSet []string
	for _, val := range counterChain.ValSet.Vals {
		valSet = append(valSet, val.Addr)
	}
	msg := ismtypes.NewMsgSetDefaultIsm(signer, []*ismtypes.OriginsMultiSigIsm{
		{
			Origin: 112,
			Ism: &ismtypes.MultiSigIsm{
				Threshold: 2,
				ValidatorPubKeys: valSet,
			},
		},
	})
	_, err := suite.ismkeeper.SetDefaultIsm(suite.ismctx, msg)
	suite.Require().NoError(err)

	sender := "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7"
	contract := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	destDomain := uint32(10)
	message, proof := counterChain.CreateMessage(sender, destDomain, contract, "Hello!")
	// Create metadata
	metadata := counterChain.CreateMetadata(message, proof)
	// Process message
	res, err := suite.msgServer.Process(suite.ctx, &types.MsgProcess{
		Sender: "cosmos108tuxfep73qjru5qsf5wxpz5hmujg53lyq0edg",
		Metadata: hexutil.Encode(metadata),
		Message: hexutil.Encode(message),
	})
	suite.Require().NoError(err)

	fmt.Println("Res: ", res)

}*/
