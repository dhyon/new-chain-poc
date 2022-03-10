package dex

const GOCTX_KEY = "dex_goctx"

type LimitOrder struct {
	Price    int64
	Quantity int64
	Creator  string
	Long     bool
}

type MarketOrder struct {
	Quantity int64
	Creator  string
	Long     bool
}

type Orders struct {
	LimitBuys   []LimitOrder
	LimitSells  []LimitOrder
	MarketBuys  []MarketOrder
	MarketSells []MarketOrder
}

func NewOrders() Orders {
	return Orders{
		LimitBuys:   []LimitOrder{},
		LimitSells:  []LimitOrder{},
		MarketBuys:  []MarketOrder{},
		MarketSells: []MarketOrder{},
	}
}
