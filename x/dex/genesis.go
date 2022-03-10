package dex

import (
	"github.com/codchen/new-chain-poc/x/dex/keeper"
	"github.com/codchen/new-chain-poc/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the longBook
	for _, elem := range genState.LongBookList {
		k.SetLongBook(ctx, elem)
	}

	// Set longBook count
	k.SetLongBookCount(ctx, genState.LongBookCount)
	// Set all the shortBook
	for _, elem := range genState.ShortBookList {
		k.SetShortBook(ctx, elem)
	}

	// Set shortBook count
	k.SetShortBookCount(ctx, genState.ShortBookCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.LongBookList = k.GetAllLongBook(ctx)
	genesis.LongBookCount = k.GetLongBookCount(ctx)
	genesis.ShortBookList = k.GetAllShortBook(ctx)
	genesis.ShortBookCount = k.GetShortBookCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
