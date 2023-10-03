package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	idMap := make([]string, 128)
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
	suite.Require().Equal(8, countPopulatedSlices(suite.keeper.Branch)) // 2^7 + 36 = 100  .. only 8 levels will be populated.

	// Adding delivered message ids to the exported state
	for i := 0; i < 100; i++ {
		gs.DeliveredMessages = append(gs.DeliveredMessages, &types.MessageDelivered{
			Id: idMap[i],
		})
	}

	// Resetting and logging
	suite.SetupTest()

	// Importing state
	err := suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	suite.Require().Equal(8, countPopulatedSlices(suite.keeper.Tree.Branch))
	suite.Require().Equal(100, len(suite.keeper.Delivered))
}

func countPopulatedSlices(arr [32][]byte) int {
	fmt.Printf("%+v", arr)
	count := 0
	for _, slice := range arr {
		if len(slice) > 0 {
			count++
		}
	}
	return count
}
