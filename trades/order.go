package trades

const (
	_ = iota
	Limit
	Stop
)

const (
	_ = iota
	pBid
	pBidPlus
	pBidMid
	pMid
	pMidPlus
	pMidMid
	pAskMinus
	pAsk
	pAskPlus
	pAskDouble
	pAskPercent
)

type Order struct {
	Symbol
	Type   uint
	Amount int
	Price  float64
}

func (o *Order) SetPrice(priceType uint) {
	var price float64
	dir := 1

	bid, ask := o.Symbol.Bid, o.Symbol.Ask

	if o.Amount > 0 {
		price = bid
	} else {
		price = ask
		dir = -1
	}

	switch priceType {
	case pBid:
		o.Price = price
	case pBidPlus:
		o.Price = price + dir*o.Tick
	case pBidMid:
		o.Price = price + dir*(ask-bid)/4
	case pMid:
		o.Price = price + dir*(ask-bid)/2
	case pMidPlus:
		o.Price = price + dir*((ask-bid)/2+o.Tick)
	case pMidMid:
		o.Price = price + dir*3*(ask-bid)/4
	case pAskMinus:
		o.Price = price + dir*(ask-bid) - o.Tick
	case pAsk:
		o.Price = bid + ask - price
	case pAskPlus:
		o.Price = bid + ask - price + dir*o.Tick
	case pAskDouble:
		o.Price = price + dir*2*(ask-bid)
	case pAskPercent:
		o.Price = price + dir*bid/100
	}

}

// Orders

func (all []Order) Symbol(s Symbol) (orders []Order) {
	for _, one := range all {
		if one.Symbol == s {
			append(orders, one)
		}
	}
	return
}

func (all []Order) Type(t uint) (orders []Order) {
	for _, one := range all {
		if one.Type == t {
			append(orders, one)
		}
	}
	return
}

func (all []Order) Dir(d bool) (orders []Order) {
	for _, one := range all {
		if (d && one.Amount > 0) || (!d && one.Amount < 0) {
			append(orders, one)
		}
	}
	return
}
