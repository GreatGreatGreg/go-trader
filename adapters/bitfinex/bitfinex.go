package bitfinex

import (
	"log"
	"os"
	"sync"

	api "github.com/bitfinexcom/bitfinex-api-go"
	"github.com/santacruz123/go-trader/platform"
	"github.com/santacruz123/go-trader/trades"
)

var once sync.Once

var bitfinexKey string
var bitfinexSecret string

var bfxPlatform *bitfinex

type bitfinex struct {
	mu      sync.Mutex
	symbols []*trades.Symbol
	client  *api.Client
}

// Get platform
func Get() platform.Platformer {
	once.Do(func() {

		bfxPlatform = &bitfinex{}

		bfxPlatform.client = api.NewClient().Auth(bitfinexKey, bitfinexSecret)

		err := bfxPlatform.client.WebSocket.Connect()

		bfxPlatform.Symbol("BTCUSD")
		bfxPlatform.Symbol("LTCUSD")

		go bfxPlatform.client.WebSocket.Subscribe()

		if err != nil {
			log.Fatal("Error connecting to bitfinex socket")
		}
	})

	return bfxPlatform
}

func (platform *bitfinex) ClosePlatform() {
	platform.client.WebSocket.Close()
}

func (platform *bitfinex) symbol(s string) *trades.Symbol {
	prices := make(chan float64)
	apiPrices := make(chan []float64)

	platform.client.WebSocket.AddSubscribe(api.CHAN_BOOK, s, apiPrices)

	go func() {
		for pack := range apiPrices {
			select {
			case prices <- pack[0]:
			default:
			}
		}

		log.Println("Bitfinex", s, "channel died")
		close(prices)
	}()

	return trades.NewSymbol(s, trades.Fx, 0.01, prices)
}

func (platform *bitfinex) Symbol(s string) (symbol *trades.Symbol, err error) {
	platform.mu.Lock()
	defer platform.mu.Unlock()

	for _, one := range platform.symbols {
		if one.Symbol() == s {
			return one, nil
		}
	}

	symbol = platform.symbol(s)
	platform.symbols = append(platform.symbols, symbol)

	return
}

func (platform *bitfinex) Orders() (orders trades.Orders, err error) {
	return
}

func (platform *bitfinex) Positions() (positions trades.Positions, err error) {
	return
}

func (platform *bitfinex) Order(o trades.Order) (id uint, err error) {
	return
}

func (platform *bitfinex) Cancel(id uint) (err error) {
	return
}

func (platform *bitfinex) Modify(id uint, order trades.Order) (err error) {
	return
}

func init() {

	bitfinexKey = os.Getenv("bitfinex_key")
	bitfinexSecret = os.Getenv("bitfinex_secret")

	if bitfinexKey == "" {
		log.Fatalf("Missing bitfinex_key")
	}

	if bitfinexSecret == "" {
		log.Fatalf("Missing bitfinex_secret")
	}
}
