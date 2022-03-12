package dex

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/client/cli"
	"github.com/codchen/new-chain-poc/x/dex/exchange"
	"github.com/codchen/new-chain-poc/x/dex/keeper"
	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the capability module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns the capability module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the capability module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the capability module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the capability module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Name returns the capability module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// Route returns the capability module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

// QuerierRoute returns the capability module's query routing key.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// LegacyQuerierHandler returns the capability module's Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the capability module's invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the capability module's genesis initialization It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion implements ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 2 }

// BeginBlock executes all ABCI BeginBlock logic respective to the capability module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock executes all ABCI EndBlock logic respective to the capability module. It
// returns no validator updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	orders := ctx.Context().Value(dexcache.GOCTX_KEY).(dexcache.Orders)
	allExistingBuys := []types.OrderBook{}
	for _, lb := range am.keeper.GetAllLongBook(ctx) {
		allExistingBuys = append(allExistingBuys, &lb)
	}
	allExistingSells := []types.OrderBook{}
	for _, sb := range am.keeper.GetAllShortBook(ctx) {
		allExistingSells = append(allExistingSells, &sb)
	}
	_, mktBuyDirty, executedBuyQuantities := exchange.MatchMarketOrders(
		orders.MarketBuys,
		allExistingSells,
	)
	_, mktSellDirty, executedSellQuantities := exchange.MatchMarketOrders(
		orders.MarketSells,
		allExistingBuys,
	)
	_, limitBuyDirty, limitSellDirty := exchange.MatchLimitOrders(
		orders.LimitBuys,
		orders.LimitSells,
		allExistingBuys,
		allExistingSells,
	)
	idToBuys := map[uint64]types.OrderBook{}
	idToSells := map[uint64]types.OrderBook{}
	for _, buy := range allExistingBuys {
		idToBuys[buy.GetId()] = buy
	}
	for _, sell := range allExistingSells {
		idToSells[sell.GetId()] = sell
	}

	for i, quantity := range executedBuyQuantities {
		if quantity > 0 {
			// TODO: settlement & notification to market order creators
			_, _ = i, quantity
		}
	}
	for i, quantity := range executedSellQuantities {
		if quantity > 0 {
			// TODO: settlement & notification to market order creators
			_, _ = i, quantity
		}
	}

	dirtyBuys := map[uint64]bool{}
	dirtySells := map[uint64]bool{}
	for id := range mktBuyDirty {
		dirtySells[id] = true
	}
	for id := range mktSellDirty {
		dirtyBuys[id] = true
	}
	for id := range limitBuyDirty {
		dirtyBuys[id] = true
	}
	for id := range limitSellDirty {
		dirtySells[id] = true
	}

	for id := range dirtyBuys {
		order := idToBuys[id]
		creatorToExecuted := exchange.RebalanceAllocations(order)
		// TODO: settlement & notification
		_ = creatorToExecuted

		if order.GetEntry().Quantity == 0 {
			am.keeper.RemoveLongBookByPrice(ctx, uint32(order.GetEntry().Price))
		} else {
			longOrder := order.(*types.LongBook)
			am.keeper.SetLongBook(ctx, *longOrder)
		}
	}
	for id := range dirtySells {
		order := idToSells[id]
		creatorToExecuted := exchange.RebalanceAllocations(order)
		// TODO: settlement & notification
		_ = creatorToExecuted

		if order.GetEntry().Quantity == 0 {
			am.keeper.RemoveShortBookByPrice(ctx, uint32(order.GetEntry().Price))
		} else {
			shortOrder := order.(*types.ShortBook)
			am.keeper.SetShortBook(ctx, *shortOrder)
		}
	}
	return []abci.ValidatorUpdate{}
}
