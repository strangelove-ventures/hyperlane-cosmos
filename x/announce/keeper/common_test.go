package keeper_test

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/mock/gomock"

	testutil "github.com/strangelove-ventures/hyperlane-cosmos/x/announce/testutil"
)

// init mocks of expected keepers that the announce module depends on
func initMocks(mailboxKeeper *testutil.MockMailboxKeeper) {
	mailboxAddress, _ := hex.DecodeString("000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3")
	domain := uint32(23456)

	mailboxKeeper.EXPECT().GetMailboxAddress().DoAndReturn(func() []byte {
		return mailboxAddress
	}).AnyTimes()
	mailboxKeeper.EXPECT().GetDomain(gomock.Any()).DoAndReturn(func(_ sdk.Context) uint32 {
		return domain
	}).AnyTimes()
}
