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

func createNShortBook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ShortBook {
	items := make([]types.ShortBook, n)
	for i := range items {
		items[i].Id = keeper.AppendShortBook(ctx, items[i])
	}
	return items
}

func TestShortBookGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShortBook(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetShortBook(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestShortBookRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShortBook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShortBook(ctx, item.Id)
		_, found := keeper.GetShortBook(ctx, item.Id)
		require.False(t, found)
	}
}

func TestShortBookGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShortBook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShortBook(ctx)),
	)
}

func TestShortBookCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNShortBook(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetShortBookCount(ctx))
}
