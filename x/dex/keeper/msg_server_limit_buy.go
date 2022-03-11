package keeper

import (
	"context"

	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) LimitBuy(goCtx context.Context, msg *types.MsgLimitBuy) (*types.MsgLimitBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	orders := ctx.Context().Value(dexcache.GOCTX_KEY).(dexcache.Orders)
	orders.LimitBuys = append(orders.LimitBuys, dexcache.LimitOrder{
		Creator:  msg.Creator,
		Price:    int64(msg.Price),
		Quantity: int64(msg.Quantity),
		Long:     true,
	})

	return &types.MsgLimitBuyResponse{}, nil
}
