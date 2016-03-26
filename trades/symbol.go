package trades

import (
	"fmt"
	"log"
	"sync"
)

// AssetType - type of the asset
type AssetType uint8

const (
	_ = iota

	// Fx - Forex
	Fx AssetType = iota

	// Stock - Stock
	Stock

	// Cfd - CFD
	Cfd

	// Option - Option
	Option
)

// Quotes structure
type Quotes struct {
	Bid float64
	Ask float64
}

// Symbol - structure for symbols
type Symbol struct {
	mu       sync.Mutex
	symbol   string
	t        AssetType
	tickSize float64
	bid      float64
	ask      float64
	subs     []chan<- Quotes
}

func (s *Symbol) String() string {
	return fmt.Sprintf("Symbol: %s bid:%.2f ask:%.2f", s.symbol, s.bid, s.ask)
}

// Symbol - returns symbol as string
func (s *Symbol) Symbol() string {
	return s.symbol
}

// TickSize - returns tickSize
func (s *Symbol) TickSize() float64 {
	return s.tickSize
}

// Bid - returns bid
func (s *Symbol) Bid() float64 {
	s.mu.Lock()
	val := s.bid
	s.mu.Unlock()
	return val
}

// Ask - returns ask
func (s *Symbol) Ask() float64 {
	s.mu.Lock()
	val := s.ask
	s.mu.Unlock()
	return val
}

// Sub - Subcribes to both bid and ask channels
func (s *Symbol) Sub(ch chan<- Quotes) {
	s.subs = append(s.subs, ch)
}

func (s *Symbol) pub(quote Quotes) {
	for _, ch := range s.subs {
		select {
		case ch <- quote:
		default:
		}
	}
}

// NewSymbol - creates new symbol
func NewSymbol(
	s string,
	t AssetType,
	size float64,
	quotes <-chan Quotes,
) (symbol *Symbol) {

	symbol = &Symbol{
		symbol:   s,
		t:        t,
		tickSize: size,
	}

	go func() {
		for quote := range quotes {

			symbol.mu.Lock()
			symbol.bid = quote.Bid
			symbol.ask = quote.Ask
			symbol.mu.Unlock()

			symbol.pub(quote)
		}
		log.Println(s, "Quotes channel died")
	}()

	return
}
