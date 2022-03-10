package dex_test

import (
	"testing"

	keepertest "github.com/codchen/new-chain-poc/testutil/keeper"
	"github.com/codchen/new-chain-poc/testutil/nullify"
	"github.com/codchen/new-chain-poc/x/dex"
	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LongBookList: []types.LongBook{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LongBookCount: 2,
		ShortBookList: []types.ShortBook{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ShortBookCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LongBookList, got.LongBookList)
	require.Equal(t, genesisState.LongBookCount, got.LongBookCount)
	require.ElementsMatch(t, genesisState.ShortBookList, got.ShortBookList)
	require.Equal(t, genesisState.ShortBookCount, got.ShortBookCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
