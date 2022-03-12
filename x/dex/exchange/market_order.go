package exchange

import (
	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
)

func MatchMarketOrders(
	marketOrders []dexcache.MarketOrder,
	orderBook []types.OrderBook,
) (int32, map[uint64]int32, []int32) {
	var totalExecuted, totalPrice int32 = 0, 0
	idToExecuted := map[uint64]int32{}
	executedQuantities := []int32{}
	for _, marketOrder := range marketOrders {
		executedQuantities = append(executedQuantities, 0)
		for _, existingOrder := range orderBook {
			var executed int32
			if marketOrder.Quantity <= int64(existingOrder.GetEntry().Quantity) {
				executed = int32(marketOrder.Quantity)
			} else {
				executed = existingOrder.GetEntry().Quantity
			}
			marketOrder.Quantity -= int64(executed)
			existingOrder.GetEntry().Quantity -= executed
			totalExecuted += executed
			totalPrice += executed * int32(existingOrder.GetEntry().Price)
			if _, ok := idToExecuted[existingOrder.GetId()]; !ok {
				idToExecuted[existingOrder.GetId()] = 0
			}
			idToExecuted[existingOrder.GetId()] += executed
			executedQuantities[len(executedQuantities)-1] += executed
			if marketOrder.Quantity == 0 {
				break
			}
		}
	}
	return totalPrice / totalExecuted, idToExecuted, executedQuantities
}
