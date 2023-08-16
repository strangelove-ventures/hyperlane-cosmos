package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtestutil "github.com/cosmos/cosmos-sdk/x/gov/testutil"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/golang/mock/gomock"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"

	"github.com/stretchr/testify/suite"
)

var (
	govAcct  = authtypes.NewModuleAddress(types.ModuleName)
	distAcct = authtypes.NewModuleAddress(disttypes.ModuleName)
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      keeper.Keeper
	queryClient types.QueryClient
	msgServer   types.MsgServer
	encCfg      moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest(t *testing.T) {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()
	banktypes.RegisterInterfaces(encCfg.InterfaceRegistry)

	// gomock initializations
	ctrl := gomock.NewController(t)
	acctKeeper := govtestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := govtestutil.NewMockBankKeeper(ctrl)
	stakingKeeper := govtestutil.NewMockStakingKeeper(ctrl)
	acctKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(govAcct).AnyTimes()
	acctKeeper.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName).Return(authtypes.NewEmptyModuleAccount(types.ModuleName)).AnyTimes()
	trackMockBalances(bankKeeper)

	stakingKeeper.EXPECT().BondDenom(ctx).Return("stake").AnyTimes()
	acctKeeper.EXPECT().GetModuleAddress(types.ModuleName).Return(govAcct).AnyTimes()
	acctKeeper.EXPECT().GetModuleAddress(disttypes.ModuleName).Return(distAcct).AnyTimes()
	acctKeeper.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName).Return(authtypes.NewEmptyModuleAccount(types.ModuleName)).AnyTimes()

	suite.ctx = ctx
	suite.keeper = keeper.NewKeeper(
		encCfg.Codec,
		key,
		bankKeeper,
		stakingKeeper,
		"",
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
