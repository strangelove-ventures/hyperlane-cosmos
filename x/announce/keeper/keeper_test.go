package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
	mbKeeper "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/keeper"
	mbtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"

	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	keeper        keeper.Keeper
	mailboxKeeper mbKeeper.ReadOnlyMailboxKeeper
	domain        uint32
	queryClient   types.QueryClient
	msgServer     types.MsgServer
	encCfg        moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest(t *testing.T) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	mailboxKey := sdk.NewKVStoreKey(mbtypes.StoreKey)

	testCtx := testutil.DefaultContextWithDB(suite.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()
	mailboxKeeper := mbKeeper.NewKeeper(encCfg.Codec, mailboxKey, nil, nil)
	mailboxKeeper.SetDomain(ctx, 23456)
	suite.domain = 23456
	suite.mailboxKeeper = mailboxKeeper

	suite.ctx = ctx
	suite.keeper = keeper.NewKeeper(
		encCfg.Codec,
		key,
		mailboxKeeper,
	)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, suite.keeper)
	queryClient := types.NewQueryClient(queryHelper)
	msgServer := keeper.NewMsgServerImpl(suite.keeper)

	suite.queryClient = queryClient
	suite.msgServer = msgServer
	suite.encCfg = encCfg
}
