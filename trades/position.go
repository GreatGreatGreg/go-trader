package trades

// Position - type responsible for position
type Position struct {
	*Symbol
	Amount float64
	Price  float64
}

// Positions - slice of positions
type Positions []Position

// Symbol - filters positions by symbol
func (all Positions) Symbol(s string) (positions Positions) {
	for _, one := range all {
		if one.Symbol.Symbol() == s {
			positions = append(positions, one)
		}
	}
	return
}

// Dir - filters positions by direction - true = buy, false = sell
func (all Positions) Dir(d bool) (positions Positions) {
	for _, one := range all {
		if (d && one.Amount > 0) || (!d && one.Amount < 0) {
			positions = append(positions, one)
		}
	}
	return
}

// PnL - gets position's PnL
func (p Position) PnL() (pnl float64) {

	if p.Amount > 0 {
		pnl = float64(p.Amount) * (p.Symbol.Bid() - p.Price)
	} else {
		pnl = float64(p.Amount) * (p.Price - p.Symbol.Ask())
	}

	return
}

// Close - creates order to close position
func (p Position) Close(t OrderPriceType) (order Order) {

	order = Order{
		Symbol: p.Symbol,
		Amount: -p.Amount,
	}

	order.FastPrice(t)
	return
}

// Scale - scales position
func (p Position) Scale(percent float64, t OrderPriceType) (order Order) {

	order = Order{
		Symbol: p.Symbol,
		Amount: p.Amount * (percent - 1),
	}

	order.FastPrice(t)
	return
}
