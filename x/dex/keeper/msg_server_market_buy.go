package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
)

func (k msgServer) MarketBuy(goCtx context.Context, msg *types.MsgMarketBuy) (*types.MsgMarketBuyResponse, error) {
	k.Orders.MarketBuys = append(k.Orders.MarketBuys, dexcache.MarketOrder{
		Creator:  msg.Creator,
		Quantity: int64(msg.Quantity),
		Long:     true,
	})

	return &types.MsgMarketBuyResponse{}, nil
}
