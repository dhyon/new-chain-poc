package exchange

import "github.com/codchen/new-chain-poc/x/dex/types"

func RebalanceAllocations(order types.OrderBook) map[string]int32 {
	newTotal := order.GetEntry().Quantity
	var oldTotal int32 = 0
	for _, allo := range order.GetEntry().GetAllocation() {
		oldTotal += allo
	}
	res := map[string]int32{}
	var acc int32 = 0
	for i, creator := range order.GetEntry().AllocationCreator {
		res[creator] = order.GetEntry().GetAllocation()[i] * newTotal / oldTotal
		acc += res[creator]
	}
	numCreators := len(order.GetEntry().AllocationCreator)
	for i := acc; i < newTotal; i++ {
		res[order.GetEntry().AllocationCreator[(i-acc)%int32(numCreators)]] += 1
	}
	return res
}
