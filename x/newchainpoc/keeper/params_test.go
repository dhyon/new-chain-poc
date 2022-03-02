package keeper_test

import (
	"testing"

	testkeeper "github.com/codchen/new-chain-poc/testutil/keeper"
	"github.com/codchen/new-chain-poc/x/newchainpoc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewchainpocKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
