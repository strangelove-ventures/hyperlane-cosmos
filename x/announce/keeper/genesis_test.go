package keeper_test

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	var err error
	valPrivKey := "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61" // Testing only, do NOT use
	storageLocation := "file:///tmp//signatures-simd"
	numMsgs := 5

	mailbox := suite.mailboxKeeper.GetMailboxAddress()
	announcementMap := map[string][]string{}

	for i := 0; i < numMsgs; i++ {
		msg, _, err := suite.mockAnnounce(valPrivKey, fmt.Sprintf("%s%d", storageLocation, i), mailbox)
		suite.Require().NoError(err)
		valAddrHex := hexutil.Encode(msg.Validator)
		announcementMap[valAddrHex] = append(announcementMap[valAddrHex], msg.StorageLocation)
	}

	// Export genesis
	gs := suite.keeper.ExportGenesis(suite.ctx)

	for _, gsAnnouncement := range gs.Announcements {
		valHex := hexutil.Encode(gsAnnouncement.Validator)
		expectedAnnouncements, ok := announcementMap[valHex]
		suite.Require().True(ok)
		for _, curr := range gsAnnouncement.Announcements.Announcement {
			suite.Require().Contains(expectedAnnouncements, curr.StorageLocation)
		}
	}

	// Resetting suite
	suite.SetupTest(suite.T())

	// Importing state
	err = suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	announcedAddresses := [][]byte{}

	for val := range announcementMap {
		announcedAddresses = append(announcedAddresses, hexutil.MustDecode(val))
	}

	// Format the announced storage locations request
	announcedLocationsReq := &types.GetAnnouncedStorageLocationsRequest{
		Validator: [][]byte{},
	}
	announcedLocationsReq.Validator = append(announcedLocationsReq.Validator, announcedAddresses...)

	// Comparing the two states to make sure they are the same
	announcedLocationsResp, err := suite.keeper.GetAnnouncedStorageLocations(suite.ctx, announcedLocationsReq)
	suite.Require().NoError(err)

	for i, announcedAddress := range announcedAddresses {
		locations, ok := announcementMap[hexutil.Encode(announcedAddress)]
		suite.Require().True(ok)

		for _, location := range locations {
			suite.Require().Contains(announcedLocationsResp.Metadata[i].Metadata, location)
		}
	}
}
