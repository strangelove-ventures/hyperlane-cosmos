package keeper_test

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	var err error
	valPrivKey := "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61" // Testing only, do NOT use
	storageLocation1 := "file:///tmp//signatures-simd1"
	storageLocation2 := "file:///tmp//signatures-simd2"

	mailbox := suite.mailboxKeeper.GetMailboxAddress()
	msg, _, err := suite.mockAnnounce(valPrivKey, storageLocation1, mailbox)
	suite.Require().NoError(err)
	msg2, _, err := suite.mockAnnounce(valPrivKey, storageLocation2, mailbox)
	suite.Require().NoError(err)
	announcementMap := map[int]*types.MsgAnnouncement{}
	announcementMap[0] = msg
	announcementMap[1] = msg2

	// Export genesis
	gs := suite.keeper.ExportGenesis(suite.ctx)

	for i, announcement := range gs.Announcements {
		suite.Require().Equal(announcementMap[i].Validator, announcement.Validator)
		suite.Require().Equal(announcementMap[i].StorageLocation, announcement.Announcement.StorageLocation)
	}

	// Resetting suite
	suite.SetupTest(suite.T())

	// Importing state
	err = suite.keeper.InitGenesis(suite.ctx, gs)
	suite.Require().NoError(err)

	announcedAddresses := [][]byte{}
	// key is the hex encoded bytes of the validator address, value is []storageLocations
	expectedStorageLocations := map[string][]string{}

	for _, announcement := range announcementMap {
		inArr := false
		for _, curr := range announcedAddresses {
			if bytes.Equal(curr, announcement.Validator) {
				inArr = true
			}
		}

		if !inArr {
			announcedAddresses = append(announcedAddresses, announcement.Validator)
			expectedStorageLocations[hexutil.Encode(announcement.Validator)] = []string{announcement.StorageLocation}
		} else {
			expectedStorageLocations[hexutil.Encode(announcement.Validator)] = append(expectedStorageLocations[hexutil.Encode(announcement.Validator)], announcement.StorageLocation)
		}
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
		locations := expectedStorageLocations[hexutil.Encode(announcedAddress)]
		for _, location := range locations {
			suite.Require().Contains(announcedLocationsResp.Metadata[i].Metadata, location)
		}
	}
}
