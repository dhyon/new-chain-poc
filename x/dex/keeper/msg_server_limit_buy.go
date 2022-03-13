package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
)

func (k msgServer) LimitBuy(goCtx context.Context, msg *types.MsgLimitBuy) (*types.MsgLimitBuyResponse, error) {
	k.Orders.LimitBuys = append(k.Orders.LimitBuys, dexcache.LimitOrder{
		Creator:  msg.Creator,
		Price:    int64(msg.Price),
		Quantity: int64(msg.Quantity),
		Long:     true,
	})

	return &types.MsgLimitBuyResponse{}, nil
}
