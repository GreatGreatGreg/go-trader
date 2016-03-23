package bitfinex_test

import (
	"github.com/santacruz123/go-trader/adapters/bitfinex"
	"github.com/santacruz123/go-trader/platform"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Symbol", func() {

	var platform platform.Platformer
	platform = bitfinex.Get()

	Context("BTCUSD", func() {
		It("Making", func() {
			btcusd, _ := platform.Symbol("BTCUSD")
			ltcusd, _ := platform.Symbol("LTCUSD")

			btcPrices := make(chan float64)
			ltcPrices := make(chan float64)
			btcusd.Sub(btcPrices)
			ltcusd.Sub(ltcPrices)

			Expect(<-btcPrices > 1.).Should(BeTrue())
			Expect(<-ltcPrices > 1.).Should(BeTrue())
		})
	})
})
