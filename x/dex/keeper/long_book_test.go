package keeper_test

import (
	"testing"

	keepertest "github.com/codchen/new-chain-poc/testutil/keeper"
	"github.com/codchen/new-chain-poc/testutil/nullify"
	"github.com/codchen/new-chain-poc/x/dex/keeper"
	"github.com/codchen/new-chain-poc/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNLongBook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LongBook {
	items := make([]types.LongBook, n)
	for i := range items {
		items[i].Id = keeper.AppendLongBook(ctx, items[i])
	}
	return items
}

func TestLongBookGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLongBook(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLongBook(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLongBookRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLongBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLongBook(ctx, item.Id)
		_, found := keeper.GetLongBook(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLongBookGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLongBook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLongBook(ctx)),
	)
}

func TestLongBookCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLongBook(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLongBookCount(ctx))
}
