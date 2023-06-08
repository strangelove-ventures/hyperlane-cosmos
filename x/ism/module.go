package ism

import (
	"context"
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/client/cli"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/keeper"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
)

var (
	_ module.AppModule      = (*AppModule)(nil)
	_ module.AppModuleBasic = (*AppModule)(nil)
)

// AppModuleBasic defines the basic application module used by the hyperlane mailbox.
// Only the RegisterInterfaces function needs to be implemented. All other function perform
// a no-op.
type AppModuleBasic struct{}

// Name returns the hyperlane ISM module name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec performs a no-op. Hyperlane ISM does not support amino.
func (AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

// RegisterInterfaces registers module concrete types into protobuf Any.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns an empty state
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&types.GenesisState{
		DefaultIsm: []types.Ism{},
	})
}

// ValidateGenesis performs a no-op.
func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

// RegisterGRPCGatewayRoutes registers hyperlane mailbox's routes
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		panic(err)
	}
}

// GetTxCmd implements AppModuleBasic interface
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd implements AppModuleBasic interface
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// AppModule represents the AppModule for this module
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new hyperlane mailbox module
func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		keeper: k,
	}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	var gs types.GenesisState
	err := cdc.UnmarshalJSON(bz, &gs)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal %s genesis state: %s", am.Name(), err))
	}
	err = am.keeper.InitGenesis(ctx, gs)
	if err != nil {
		panic(err)
	}
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(&gs)
}

// BeginBlock implements the AppModule interface
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// EndBlock implements the AppModule interface
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
