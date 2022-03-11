package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MarketBuy(goCtx context.Context, msg *types.MsgMarketBuy) (*types.MsgMarketBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	orders := ctx.Context().Value(dexcache.GOCTX_KEY).(dexcache.Orders)
	orders.MarketBuys = append(orders.MarketBuys, dexcache.MarketOrder{
		Creator:  msg.Creator,
		Quantity: int64(msg.Quantity),
		Long:     true,
	})

	return &types.MsgMarketBuyResponse{}, nil
}
