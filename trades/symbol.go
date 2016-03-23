package trades

import (
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

// Symbol - structure for symbols
type Symbol struct {
	mu       sync.Mutex
	symbol   string
	t        AssetType
	tickSize float64
	bid      float64
	ask      float64
	bidCh    []chan<- float64
	askCh    []chan<- float64
	priceCh  []chan<- float64
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

// SubBid - Subcribes to the bid channel
func (s *Symbol) SubBid(ch chan<- float64) {
	s.bidCh = append(s.bidCh, ch)
}

// SubAsk - Subcribes to the ask channel
func (s *Symbol) SubAsk(ch chan<- float64) {
	s.askCh = append(s.askCh, ch)
}

// Sub - Subcribes to both bid and ask channels
func (s *Symbol) Sub(ch chan<- float64) {
	s.bidCh = append(s.bidCh, ch)
	s.askCh = append(s.askCh, ch)
}

func (s *Symbol) pub(price float64, bidask bool) {
	for _, ch := range s.bidCh {
		select {
		case ch <- price:
		default:
		}
	}
}

// NewSymbol - creates new symbol
func NewSymbol(
	s string,
	t AssetType,
	size float64,
	bids <-chan float64,
	asks <-chan float64,
) (symbol *Symbol) {

	symbol = &Symbol{
		symbol:   s,
		t:        t,
		tickSize: size,
	}

	go func() {
		for {
			select {
			case bid := <-bids:
				symbol.bid = bid
				symbol.pub(bid, false)
			case ask := <-asks:
				symbol.ask = ask
				symbol.pub(ask, false)
			}
		}
	}()

	return
}
