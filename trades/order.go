package trades

// OrderPriceType - possible p
type OrderPriceType uint8

// Fast price calculation types
const (
	_                     = iota
	OptBid OrderPriceType = iota
	OptBidPlus
	OptBidMid
	OptMid
	OptMidPlus
	OptMidMid
	OptAskMinus
	OptAsk
	OptAskPlus
	OptAskDouble
	OptAskPercent
)

// Order - order structure
type Order struct {
	Symbol *Symbol
	IsStop bool
	Amount float64
	Price  float64
}

// Orders - collection of orders
type Orders []Order

// NewOrder - creates new empty order
func NewOrder() *Order {
	return &Order{}
}

// FastPrice - setting price of the order
func (o *Order) FastPrice(priceType OrderPriceType) {
	var price float64
	dir := 1.

	bid, ask := o.Symbol.Bid(), o.Symbol.Ask()

	if o.Amount > 0 {
		price = bid
	} else {
		price = ask
		dir = -1
	}

	switch priceType {
	case OptBid:
		o.Price = price
	case OptBidPlus:
		o.Price = price + dir*o.Symbol.TickSize()
	case OptBidMid:
		o.Price = price + dir*(ask-bid)/4
	case OptMid:
		o.Price = price + dir*(ask-bid)/2
	case OptMidPlus:
		o.Price = price + dir*((ask-bid)/2+o.Symbol.TickSize())
	case OptMidMid:
		o.Price = price + dir*3*(ask-bid)/4
	case OptAskMinus:
		o.Price = price + dir*(ask-bid) - o.Symbol.TickSize()
	case OptAsk:
		o.Price = bid + ask - price
	case OptAskPlus:
		o.Price = bid + ask - price + dir*o.Symbol.TickSize()
	case OptAskDouble:
		o.Price = price + dir*2*(ask-bid)
	case OptAskPercent:
		o.Price = price + dir*bid/100
	}

}

// Symbol - filters orders by symbol
func (all Orders) Symbol(s string) (orders Orders) {
	for _, one := range all {
		if one.Symbol.Symbol() == s {
			orders = append(orders, one)
		}
	}
	return
}

// LimitStop - filters orders by order type - limit = true, stop = false
func (all Orders) LimitStop(t bool) (orders Orders) {
	for _, one := range all {
		if (t && !one.IsStop) || (!t && one.IsStop) {
			orders = append(orders, one)
		}
	}
	return
}

// Limit - returns only limit orders
func (all Orders) Limit() (orders Orders) {
	for _, one := range all {
		if !one.IsStop {
			orders = append(orders, one)
		}
	}
	return
}

// Stop - filters everything but stop orders
func (all Orders) Stop() (orders Orders) {
	for _, one := range all {
		if one.IsStop {
			orders = append(orders, one)
		}
	}
	return
}

// LongShort - filters orders by type - true = long, false = short
func (all Orders) LongShort(d bool) (orders Orders) {
	for _, one := range all {
		if (d && one.Amount > 0) || (!d && one.Amount < 0) {
			orders = append(orders, one)
		}
	}
	return
}

// Long - filters everything but long orders
func (all Orders) Long() (orders Orders) {
	for _, one := range all {
		if one.Amount > 0 {
			orders = append(orders, one)
		}
	}
	return
}

// Short - filters everything but short orders
func (all Orders) Short() (orders Orders) {
	for _, one := range all {
		if one.Amount < 0 {
			orders = append(orders, one)
		}
	}
	return
}
