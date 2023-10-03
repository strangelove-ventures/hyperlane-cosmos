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

		fmt.Println(msg)

		res, err := suite.msgServer.Dispatch(suite.ctx, msg)
		suite.Require().NoError(err)
		idMap[i] = res.MessageId
		fmt.Println(idMap[i])

	}

	// Log Keeper State
	//fmt.Printf("Keeper State: %+v\n", suite.keeper)

	// Exporting Genesis and logging the length of Branches
	gs := suite.keeper.ExportGenesis(suite.ctx)
	fmt.Printf("Length of Keeper Branches: %d\n", len(suite.keeper.Branch))
	fmt.Printf("Length of Populated Branches: %d\n", countPopulatedSlices(suite.keeper.Branch))
	suite.Require().Equal(7, countPopulatedSlices(suite.keeper.Branch)) // 2^7 = 96 .. only 7 levels will be populated.

	// count := 0
	// for i, branch := range suite.keeper.Branch {
	// 	encodedBranch := hexutil.Encode(branch)
	// 	if encodedBranch == idMap[i] {
	// 		count++
	// 	}
	// }

	// //Log the count before assertion
	// fmt.Printf("Count before assertion: %d\n", count)
	// suite.Require().Equal(100, count)

	// Adding delivered message ids to the exported state
	for i := 0; i < 100; i++ {
		gs.DeliveredMessages = append(gs.DeliveredMessages, &types.MessageDelivered{
			Id: idMap[i],
		})
	}

	// Resetting and logging
	fmt.Println("Resetting and importing state...")
	suite.SetupTest()

	// Importing state
	err := suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	// Logging and checking Branches and Delivered after Import
	fmt.Printf("Length of Keeper Branches after Import: %d\n", len(suite.keeper.Branch))
	fmt.Printf("Length of Keeper Delivered after Import: %d\n", len(suite.keeper.Delivered))

	for key := range suite.keeper.Delivered {
		fmt.Println(key)
	}

	//suite.Require().Equal(uint32(100), uint32(len(suite.keeper.Branches)))
	suite.Require().Equal(100, len(suite.keeper.Delivered))
}

func countPopulatedSlices(arr [32][]byte) int {
	count := 0
	for _, slice := range arr {
		if len(slice) > 0 {
			count++
		}
	}
	return count
}
