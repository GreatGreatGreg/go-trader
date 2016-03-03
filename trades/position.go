package trades

type Position struct {
	*Symbol
	Amount int
	Price  float64
}

func (all []Position) Symbol(s string) (positions []Position) {
	for _, one := range all {
		if *one.Symbol.Symbol == s {
			append(positions, one)
		}
	}
	return
}

func (all []Position) Dir(d bool) (positions []Position) {
	for _, one := range all {
		if (d && one.Amount > 0) || (!d && one.Amount < 0) {
			append(positions, one)
		}
	}
	return
}

func (p Position) PnL() (pnl float64) {

	if p.Amount > 0 {
		pnl = p.Amount * (p.Symbol.Bid - p.Price)
	} else {
		pnl = p.Amount * (p.Price - p.Symbol.Ask)
	}

	return
}

func (p Position) Close(t uint) (order Order) {

	order = Order{
		Symbol: p.Symbol,
		Type:   t,
		Amount: -p.Amount}

	order.SetPrice(t)
	return
}

func (p Position) Scale(percent float64, t uint) (order Order) {

	order = Order{
		Symbol: p.Symbol,
		Type:   t,
		Amount: int(p.Amount * (percent - 1))}

	order.SetPrice(t)
	return
}
