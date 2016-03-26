package bitfinex_test

import (
	"github.com/santacruz123/go-trader/adapters/bitfinex"
	"github.com/santacruz123/go-trader/platform"
	"github.com/santacruz123/go-trader/trades"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order", func() {

	var platform platform.Platformer
	platform = bitfinex.Get()

	var orderId uint

	Context("Order", func() {
		It("Limit buy order", func() {

			btcusd, err := platform.Symbol("BTCUSD")
			Expect(err).Should(Succeed())

			order := trades.NewOrder()
			order.Symbol = btcusd
			order.Amount = 0.1
			order.FastPrice(trades.OptMid)

			orderId, err = platform.Order(*order)

			Expect(err).Should(Succeed())
			Expect(orderId).NotTo(BeZero())
		})

		It("Stop sell order", func() {

			btcusd, err := platform.Symbol("BTCUSD")
			Expect(err).Should(Succeed())

			order := trades.NewOrder()
			order.Symbol = btcusd
			order.Amount = -0.1
			order.Price = 1
			order.IsStop = true

			id, err := platform.Order(*order)

			Expect(err).Should(Succeed())
			Expect(id).NotTo(BeZero())
		})

		It("Cancel order", func() {
			err := platform.Cancel(orderId)
			Expect(err).Should(Succeed())
		})

		It("Cancel all orders", func() {
			err := platform.CancelAll()
			Expect(err).Should(Succeed())
		})
	})
})
