package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
)

func (k msgServer) LimitSell(goCtx context.Context, msg *types.MsgLimitSell) (*types.MsgLimitSellResponse, error) {
	k.Orders.LimitSells = append(k.Orders.LimitSells, dexcache.LimitOrder{
		Creator:  msg.Creator,
		Price:    int64(msg.Price),
		Quantity: int64(msg.Quantity),
		Long:     false,
	})

	return &types.MsgLimitSellResponse{}, nil
}
