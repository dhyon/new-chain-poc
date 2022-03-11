package dex

import (
	"math/rand"

	"github.com/codchen/new-chain-poc/testutil/sample"
	dexsimulation "github.com/codchen/new-chain-poc/x/dex/simulation"
	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = dexsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgLimitBuy = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLimitBuy int = 100

	opWeightMsgLimitSell = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLimitSell int = 100

	opWeightMsgMarketBuy = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMarketBuy int = 100

	opWeightMsgMarketSell = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMarketSell int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	dexGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&dexGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgLimitBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLimitBuy, &weightMsgLimitBuy, nil,
		func(_ *rand.Rand) {
			weightMsgLimitBuy = defaultWeightMsgLimitBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLimitBuy,
		dexsimulation.SimulateMsgLimitBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLimitSell int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLimitSell, &weightMsgLimitSell, nil,
		func(_ *rand.Rand) {
			weightMsgLimitSell = defaultWeightMsgLimitSell
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLimitSell,
		dexsimulation.SimulateMsgLimitSell(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMarketBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMarketBuy, &weightMsgMarketBuy, nil,
		func(_ *rand.Rand) {
			weightMsgMarketBuy = defaultWeightMsgMarketBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMarketBuy,
		dexsimulation.SimulateMsgMarketBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMarketSell int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMarketSell, &weightMsgMarketSell, nil,
		func(_ *rand.Rand) {
			weightMsgMarketSell = defaultWeightMsgMarketSell
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMarketSell,
		dexsimulation.SimulateMsgMarketSell(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
