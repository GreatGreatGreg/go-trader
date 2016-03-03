package trades

import "sync"

type Symbol struct {
	sync.Mutex
	Symbol string
	Tick   float64
	Bid    float64
	Ask    float64
}

type Account struct {
	Balance float64
	Margin  float64
	PnL     float64
}
