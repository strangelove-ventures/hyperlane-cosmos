package keeper_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmtime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	common "github.com/strangelove-ventures/hyperlane-cosmos/x/common"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	keeper      keeper.Keeper
	queryClient types.QueryClient
	msgServer   types.MsgServer

	encCfg moduletestutil.TestEncodingConfig
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	key := sdk.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, sdk.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	encCfg := moduletestutil.MakeTestEncodingConfig()

	suite.ctx = ctx
	suite.keeper = keeper.NewKeeper(
		encCfg.Codec,
		key,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	types.RegisterInterfaces(encCfg.InterfaceRegistry)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, suite.keeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.queryClient = queryClient
	suite.msgServer = keeper.NewMsgServerImpl(suite.keeper)
	suite.encCfg = encCfg
}

type TestCase struct {
	name     string
	message  []byte
	metadata []byte
	expPass  bool
}

func (suite *KeeperTestSuite) TestVerify() {
	var (
		msg    *types.MsgSetDefaultIsm
		signer string
	)

	testCases := []TestCase{}

	for i := 0; i < len(messages); i++ {
		message, err := hex.DecodeString(messages[i])
		suite.Require().NoError(err)
		metadata, err := hex.DecodeString(metadatas[i])
		suite.Require().NoError(err)
		testCases = append(testCases, TestCase{
			name:     fmt.Sprintf("Message %d", i),
			message:  message,
			metadata: metadata,
			expPass:  true,
		})
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			signer = authtypes.NewModuleAddress(govtypes.ModuleName).String()
			msg = types.NewMsgSetDefaultIsm(signer, defaultIsms)
			_, err := suite.msgServer.SetDefaultIsm(suite.ctx, msg)
			suite.Require().NoError(err)

			pass := suite.keeper.Verify(tc.metadata, tc.message)

			if tc.expPass {
				suite.Require().True(pass)
			} else {
				suite.Require().False(pass)
			}
		})
	}
}

func TestVerifyMerkleProof(t *testing.T) {
	for i := 0; i < len(messages); i++ {
		message, err := hex.DecodeString(messages[i])
		require.NoError(t, err)

		metadata, err := hex.DecodeString(metadatas[i])
		require.NoError(t, err)

		result := types.VerifyMerkleProof(metadata, message)
		require.True(t, result)
	}
}

func TestVerifyValidatorSignatures(t *testing.T) {
	ismMap := map[uint32]types.AbstractIsm{}
	for _, originIsm := range defaultIsms {
		ism := types.MustUnpackAbstractIsm(originIsm.AbstractIsm)
		ismMap[originIsm.Origin] = ism
	}

	for i := 0; i < len(messages); i++ {
		message, err := hex.DecodeString(messages[i])
		require.NoError(t, err)

		metadata, err := hex.DecodeString(metadatas[i])
		require.NoError(t, err)

		// Verify will also verify Merkle Proof now
		result := ismMap[common.Origin(message)].Verify(metadata, message)
		require.True(t, result)
	}
}
