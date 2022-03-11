package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MarketSell(goCtx context.Context, msg *types.MsgMarketSell) (*types.MsgMarketSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	orders := ctx.Context().Value(dexcache.GOCTX_KEY).(dexcache.Orders)
	orders.MarketSells = append(orders.MarketSells, dexcache.MarketOrder{
		Creator:  msg.Creator,
		Quantity: int64(msg.Quantity),
		Long:     false,
	})

	return &types.MsgMarketSellResponse{}, nil
}
