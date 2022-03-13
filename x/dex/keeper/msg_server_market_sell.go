package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
)

func (k msgServer) MarketSell(goCtx context.Context, msg *types.MsgMarketSell) (*types.MsgMarketSellResponse, error) {
	k.Orders.MarketSells = append(k.Orders.MarketSells, dexcache.MarketOrder{
		Creator:  msg.Creator,
		Quantity: int64(msg.Quantity),
		Long:     false,
	})

	return &types.MsgMarketSellResponse{}, nil
}
