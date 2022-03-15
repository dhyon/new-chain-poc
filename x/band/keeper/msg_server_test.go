package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/codchen/new-chain-poc/testutil/keeper"
	"github.com/codchen/new-chain-poc/x/band/keeper"
	"github.com/codchen/new-chain-poc/x/band/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BandKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
