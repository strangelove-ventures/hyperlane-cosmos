package keeper_test

import (
	"encoding/binary"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func createHyperlaneMsg(nonce uint32, origin uint32, destination uint32, senderBech32 string, recipientAddressHex string, hexMsgBody string) (hyperlaneMessageHex string) {
	// TODO: NewMessage
	var message []byte

	// TODO: Make sure this is the right version
	version := make([]byte, 1)
	message = append(message, version...)

	// Nonce is the tree count.
	nonceBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nonceBytes, nonce)
	message = append(message, nonceBytes...)

	// Local Domain is set on NewKeeper
	originBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(originBytes, origin)
	message = append(message, originBytes...)

	// Get the Sender address
	// Since this is a cosmos chain, sender will be a bech32 address
	sender := sdk.MustAccAddressFromBech32(senderBech32).Bytes()
	for len(sender) < (common.DESTINATION_OFFSET - common.SENDER_OFFSET) {
		padding := make([]byte, 1)
		sender = append(padding, sender...)
	}
	message = append(message, sender...)

	// Get the Destination Domain
	destinationBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(destinationBytes, destination)
	message = append(message, destinationBytes...)

	// Get the Recipient address
	// Since the recipient could be any destination change, the address must be in hex, non-bech32 format
	recipient := hexutil.MustDecode(recipientAddressHex)
	for len(recipient) < (common.BODY_OFFSET - common.RECIPIENT_OFFSET) {
		padding := make([]byte, 1)
		recipient = append(padding, recipient...)
	}
	message = append(message, recipient...)

	// Get the Message Body
	// messageBytes := []byte(msg.MessageBody)
	messageBytes := hexutil.MustDecode(hexMsgBody)
	if len(messageBytes) > MAX_MESSAGE_BODY_BYTES {
		panic("msg too long")
	}
	message = append(message, messageBytes...)
	hyperlaneMessageHex = hexutil.Encode(message)
	return
}

const MAX_MESSAGE_BODY_BYTES = 2_000

func (suite *KeeperTestSuite) TestDispatch() {
	var (
		msg             *types.MsgDispatch
		id              string
		hyperlaneMsgHex string
		senderHex       string
	)

	// This is set in Setuptest for the keeper test suite
	origin := uint32(10)

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
				senderBytes := sdk.MustAccAddressFromBech32(sender).Bytes()
				senderHex = hexutil.Encode(senderBytes) // no padding

				recipientBech32 := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
				recipientBytes := sdk.MustAccAddressFromBech32(recipientBech32).Bytes()
				recipientHex := hexutil.Encode(recipientBytes)
				domain := uint32(12)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(sender, domain, recipientHex, msgBody)
				hyperlaneMsgHex = createHyperlaneMsg(0, origin, domain, sender, recipientHex, msgBody)
			},
			true,
		},
		{
			"success with account as sender",
			func() {
				id = "0x0c4024fb3262d8852b1fc5caa9f73f91b9375a1cbe51f6da52d192d272623fd1"
				senderHex = "0x000000000000000000000000f7edfb284f4659b69aab64bb8f23cb28460c091a"
				sender := sdk.MustBech32ifyAddressBytes("cosmos", hexutil.MustDecode(senderHex))

				recipientBech32 := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
				recipientBytes := sdk.MustAccAddressFromBech32(recipientBech32).Bytes()
				recipientHex := hexutil.Encode(recipientBytes)
				domain := uint32(12)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(sender, domain, recipientHex, msgBody)
				hyperlaneMsgHex = createHyperlaneMsg(0, origin, domain, sender, recipientHex, msgBody)

			},
			true,
		},
		{
			"success with hyperlane explorer data w/ sender/recipient padding",
			func() {
				id = "0x82fbf348d68e34903627d1f57fa3227211e35134e2a613ce6b78ce2d42e17198"
				senderHex = "0x0000000000000000000000002b0860e52244f03e59f12cfe413d6a29bc30b893"
				senderBytes := hexutil.MustDecode(senderHex)
				senderBech32 := sdk.MustBech32ifyAddressBytes("cosmos", senderBytes)
				recipient := "0x00000000000000000000000076a2f655352752af6ce9b03932b9090009dc5d0c"
				domain := uint32(43114)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(senderBech32, domain, recipient, msgBody)
				hyperlaneMsgHex = createHyperlaneMsg(0, origin, domain, senderBech32, recipient, msgBody)
			},
			true,
		},
		{
			"success with hyperlane explorer data w/o sender/recipient padding",
			func() {
				id = "0x82fbf348d68e34903627d1f57fa3227211e35134e2a613ce6b78ce2d42e17198"
				senderHex = "0x0000000000000000000000002b0860e52244f03e59f12cfe413d6a29bc30b893"
				senderBytes := hexutil.MustDecode(senderHex)
				senderBech32 := sdk.MustBech32ifyAddressBytes("cosmos", senderBytes)
				recipient := "0x76a2f655352752af6ce9b03932b9090009dc5d0c"
				domain := uint32(43114)
				msgBody := hexutil.Encode([]byte("Hello!"))
				msg = types.NewMsgDispatch(senderBech32, domain, recipient, msgBody)
				hyperlaneMsgHex = createHyperlaneMsg(0, origin, domain, senderBech32, recipient, msgBody)
			},
			true,
		},
		{
			"fails with wrong id match",
			func() {
				id = "0x82fbf348d68e34903627d1f57fa3227211e35134e2a613ce6b78ce2d42e17198"
				senderHex = "0x0000000000000000000000002b0860e52244f03e59f12cfe413d6a29bc30b893"
				senderBytes := hexutil.MustDecode(senderHex)
				senderBech32 := sdk.MustBech32ifyAddressBytes("cosmos", senderBytes)
				recipient := "0x76a2f655352752af6ce9b03932b9090009dc5d0c"
				domain := uint32(43114)
				msgBody := hexutil.Encode([]byte("Hello")) // Removed !
				msg = types.NewMsgDispatch(senderBech32, domain, recipient, msgBody)
				hyperlaneMsgHex = createHyperlaneMsg(0, origin, domain, senderBech32, recipient, msgBody)
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

			//zero-pad the sender w/ appropriate hyperlane byte length
			senderB, _ := hexutil.Decode(senderHex)
			for len(senderB) < (common.DESTINATION_OFFSET - common.SENDER_OFFSET) {
				padding := make([]byte, 1)
				senderB = append(padding, senderB...)
			}

			// Verify events
			expectedEvents := sdk.Events{
				sdk.NewEvent(
					types.EventTypeDispatch,
					sdk.NewAttribute(types.AttributeKeyDestination, strconv.FormatUint(uint64(msg.DestinationDomain), 10)),
					sdk.NewAttribute(types.AttributeKeyHyperlaneMessage, hyperlaneMsgHex),
					sdk.NewAttribute(types.AttributeKeyMessage, msg.MessageBody),
					sdk.NewAttribute(types.AttributeKeyNonce, "0"), // TODO: Nonce is always zero
					sdk.NewAttribute(types.AttributeKeyOrigin, strconv.FormatUint(testOriginDomain, 10)),
					sdk.NewAttribute(types.AttributeKeyRecipientAddress, msg.RecipientAddress),
					sdk.NewAttribute(types.AttributeKeySender, senderHex),
					sdk.NewAttribute(types.AttributeKeyVersion, "0"), // TODO(nix): Figure out version
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
