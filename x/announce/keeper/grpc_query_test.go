package keeper_test

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func (suite *KeeperTestSuite) TestQueryAnnouncements() {
	suite.SetupTest(suite.T())
	valPrivKey := "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61" // Testing only, do NOT use
	mailbox := suite.mailboxKeeper.GetMailboxAddress()
	storageLocation := "file:///tmp//signatures-simd1"
	msg, _, err := suite.mockAnnounce(valPrivKey, storageLocation, mailbox)
	suite.Require().NoError(err)
	req := &types.GetAnnouncedValidatorsRequest{}
	resp, err := suite.queryClient.GetAnnouncedValidators(suite.ctx, req)

	suite.Require().NoError(err)
	suite.Require().Equal(hexutil.Encode(msg.Validator), resp.Validator[0])

	vals := [][]byte{}
	for _, val := range resp.Validator {
		valHex, err := hexutil.Decode(val)
		suite.Require().NoError(err)
		vals = append(vals, valHex)
	}

	storageLocReq := &types.GetAnnouncedStorageLocationsRequest{
		Validator: vals,
	}
	storageLocResp, err := suite.queryClient.GetAnnouncedStorageLocations(suite.ctx, storageLocReq)
	suite.Require().NoError(err)
	suite.Require().Equal(storageLocResp.Metadata[0].Metadata[0], msg.StorageLocation)
}
