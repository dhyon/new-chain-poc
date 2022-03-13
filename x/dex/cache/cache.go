package dex

import "fmt"

const LIMIT_BUY_EVENT_TYPE = "dex_lb"
const LIMIT_SELL_EVENT_TYPE = "dex_ls"
const MARKET_BUY_EVENT_TYPE = "dex_mb"
const MARKET_SELL_EVENT_TYPE = "dex_ms"
const CREATOR_ATTR = "creator"
const PRICE_ATTR = "price"
const QUANTITY_ATTR = "quantity"

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

func (o Orders) String() string {
	return fmt.Sprintf("Limit Buys: %d, Limit Sells: %d, Market Buys: %d, Market Sells: %d", len(o.LimitBuys), len(o.LimitSells), len(o.MarketBuys), len(o.MarketSells))
}

func NewOrders() Orders {
	return Orders{
		LimitBuys:   []LimitOrder{},
		LimitSells:  []LimitOrder{},
		MarketBuys:  []MarketOrder{},
		MarketSells: []MarketOrder{},
	}
}
