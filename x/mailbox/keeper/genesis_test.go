package keeper_test

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	idMap := make([]string, 128) // Will populate 8 levels.
	for i := 0; i < 128; i++ {
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

	// Exporting Genesis and logging the length of Branches
	gs := suite.keeper.ExportGenesis(suite.ctx)

	// Adding delivered message ids to the exported state
	for i := 0; i < 128; i++ {
		gs.DeliveredMessages = append(gs.DeliveredMessages, &types.MessageDelivered{
			Id: idMap[i],
		})
	}

	// Checking the exported state
	suite.Require().Equal(8, countGenesisStatePopulatedSlices((gs.Tree.Branch)))
	suite.Require().Equal(128, len(gs.DeliveredMessages))

	// Resetting suite
	suite.SetupTest()

	// Importing state
	err := suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	// Comparing the two states to make sure they are the same
	tree := suite.keeper.GetImtTree(suite.ctx)
	suite.Require().Equal(8, countKeeperPopulatedSlices(tree.Branch))

	// Verify each delivered message is found
	for i := 0; i < 128; i++ {
		id, err := hexutil.Decode(idMap[i])
		suite.Require().NoError(err)
		req := types.QueryMsgDeliveredRequest{
			MessageId: id,
		}
		resp, err := suite.queryClient.MsgDelivered(suite.ctx, &req)
		suite.Require().NoError(err)
		suite.Require().True(resp.Delivered, "Message delivered was not found")
	}

	// Check that the branches match
	isEqual := reflect.DeepEqual(gs.Tree.Branch, tree.Branch[:])
	suite.Require().True(isEqual, "Imported state is not the same as exported state")
}

func countKeeperPopulatedSlices(arr [32][]byte) int {
	count := 0
	for _, slice := range arr {
		if len(slice) > 0 {
			count++
		}
	}
	return count
}

func countGenesisStatePopulatedSlices(arr [][]byte) int {
	count := 0
	for _, slice := range arr {
		if len(slice) > 0 {
			count++
		}
	}
	return count
}
