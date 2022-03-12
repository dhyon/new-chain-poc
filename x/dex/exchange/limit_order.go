package exchange

import (
	dexcache "github.com/codchen/new-chain-poc/x/dex/cache"
	"github.com/codchen/new-chain-poc/x/dex/types"
)

func MatchLimitOrders(
	longOrders []dexcache.LimitOrder,
	shortOrders []dexcache.LimitOrder,
	longBook []types.OrderBook,
	shortBook []types.OrderBook,
) (int32, map[uint64]int32, map[uint64]int32) {
	longIdToExecuted := map[uint64]int32{}
	shortIdToExecuted := map[uint64]int32{}
	for _, order := range longOrders {
		addOrderToOrderBook(order, longBook, longIdToExecuted)
	}
	for _, order := range shortOrders {
		addOrderToOrderBook(order, shortBook, shortIdToExecuted)
	}
	var totalExecuted, totalPrice int32 = 0, 0
	var longPtr, shortPtr = len(longBook) - 1, 0
	for longPtr >= 0 && shortPtr < len(shortBook) && longBook[longPtr].GetEntry().Price >= shortBook[shortPtr].GetEntry().Price {
		var executed int32
		if longBook[longPtr].GetEntry().Quantity < shortBook[shortPtr].GetEntry().Quantity {
			executed = longBook[longPtr].GetEntry().Quantity
		} else {
			executed = shortBook[shortPtr].GetEntry().Quantity
		}
		longBook[longPtr].GetEntry().Quantity -= executed
		shortBook[shortPtr].GetEntry().Quantity -= executed
		if _, ok := longIdToExecuted[longBook[longPtr].GetId()]; ok {
			longIdToExecuted[longBook[longPtr].GetId()] += executed
		} else {
			longIdToExecuted[longBook[longPtr].GetId()] = executed
		}
		if _, ok := shortIdToExecuted[shortBook[shortPtr].GetId()]; ok {
			shortIdToExecuted[shortBook[shortPtr].GetId()] += executed
		} else {
			shortIdToExecuted[shortBook[shortPtr].GetId()] = executed
		}

		totalExecuted += executed * 2
		totalPrice += executed * (longBook[longPtr].GetEntry().GetPrice() + shortBook[shortPtr].GetEntry().GetPrice())

		if longBook[longPtr].GetEntry().Quantity == 0 {
			longPtr -= 1
		}
		if shortBook[shortPtr].GetEntry().Quantity == 0 {
			shortPtr += 1
		}
	}
	return totalPrice / totalExecuted, longIdToExecuted, shortIdToExecuted
}

func addOrderToOrderBook(
	order dexcache.LimitOrder,
	orderBook []types.OrderBook,
	dirty map[uint64]int32,
) {
	isLong := order.Long
	insertAt := -1
	for i, ob := range orderBook {
		if ob.GetEntry().Price == int32(order.Price) {
			if _, ok := dirty[ob.GetId()]; !ok {
				dirty[ob.GetId()] = 0
			}
			ob.GetEntry().Quantity += int32(order.Quantity)
			existing := false
			for j, allocation := range ob.GetEntry().AllocationCreator {
				if allocation == order.Creator {
					existing = true
					ob.GetEntry().Allocation[j] += int32(order.Quantity)
					break
				}
			}
			if !existing {
				ob.GetEntry().AllocationCreator = append(ob.GetEntry().AllocationCreator, order.Creator)
				ob.GetEntry().Allocation = append(ob.GetEntry().Allocation, int32(order.Quantity))
			}
			return
		}
		if isLong {
			if order.Price > int64(ob.GetEntry().GetPrice()) {
				insertAt = i
				break
			}
		} else {
			if order.Price < int64(ob.GetEntry().GetPrice()) {
				insertAt = i
				break
			}
		}
	}
	var newOrder types.OrderBook
	if isLong {
		newOrder = &types.LongBook{
			Id: uint64(order.Price),
			Entry: &types.OrderEntry{
				Price:             int32(order.Price),
				Quantity:          int32(order.Quantity),
				AllocationCreator: []string{order.Creator},
				Allocation:        []int32{int32(order.Quantity)},
			},
		}
	} else {
		newOrder = &types.ShortBook{
			Id: uint64(order.Price),
			Entry: &types.OrderEntry{
				Price:             int32(order.Price),
				Quantity:          int32(order.Quantity),
				AllocationCreator: []string{order.Creator},
				Allocation:        []int32{int32(order.Quantity)},
			},
		}
	}
	if insertAt == -1 {
		orderBook = append(orderBook, newOrder)
	} else {
		orderBook = append(orderBook, &types.LongBook{})
		copy(orderBook[insertAt+1:], orderBook[insertAt:])
		orderBook[insertAt] = newOrder
	}
	dirty[uint64(order.Price)] = 0
}
