package keeper_test

import (
	"cosmossdk.io/math"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

func (suite *KeeperTestSuite) TestQueryIgp() {
	igpCreator := "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
	igpBeneficiary := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"

	suite.SetupTest(suite.T())
	igpId, err := suite.mockCreateIgp(igpCreator, igpBeneficiary, math.ZeroInt())
	suite.Require().NoError(err)

	req := &types.GetBeneficiaryRequest{IgpId: igpId}
	resp, err := suite.queryClient.GetBeneficiary(suite.ctx, req)
	suite.Require().NoError(err)
	suite.Require().Equal(igpBeneficiary, resp.GetAddress())
}
