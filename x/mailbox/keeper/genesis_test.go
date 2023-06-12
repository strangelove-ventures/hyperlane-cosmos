package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	idMap := make([]string, 100)
	for i := 0; i < 100; i++ {
		sender := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
		recipientBech32 := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
		recipientBytes := sdk.MustAccAddressFromBech32(recipientBech32).Bytes()
		recipientHex := hexutil.Encode(recipientBytes)
		domain := uint32(12)
		msgBody := hexutil.Encode([]byte(fmt.Sprintf("Hello!%d", i)))
		msg := types.NewMsgDispatch(sender, domain, recipientHex, msgBody)
		res, err := suite.msgServer.Dispatch(suite.ctx, msg)
		suite.Require().NoError(err)
		idMap[i] = res.MessageId
	}
	// TODO: add ProcessMsg tests

	res, err := suite.queryClient.CurrentTreeMetadata(suite.ctx, &types.QueryCurrentTreeMetadataRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(uint32(100), res.Count)

	// Verify tree exported correctly
	gs := suite.keeper.ExportGenesis(suite.ctx)
	suite.Require().Equal(uint32(100), gs.Tree.Count)
	count := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if gs.Tree.TreeEntries[j].Index == uint32(i) {
				if hexutil.Encode(gs.Tree.TreeEntries[j].Message) == idMap[i] {
					count++
					break;
				}
			}
		}
	}
	suite.Require().Equal(100, count)

	// Add some delivered message ids to the exported state
	for i := 0; i < 100; i++ {
		gs.DeliveredMessages = append(gs.DeliveredMessages, &types.MessageDelivered{
			Id: idMap[i],
		})
	}

	// Reset
	suite.SetupTest()

	// Import state
	err = suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	// Check Tree and Delivered
	suite.Require().Equal(uint32(100), suite.keeper.Tree.Count())
	suite.Require().Equal(100, len(suite.keeper.Delivered))
}