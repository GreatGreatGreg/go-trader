package bitfinex_test

import (
	"github.com/santacruz123/go-trader/adapters/bitfinex"
	"github.com/santacruz123/go-trader/platform"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Positions", func() {

	var platform platform.Platformer
	platform = bitfinex.Get()

	Context("Positions", func() {
		It("Get", func() {
			positions, err := platform.Positions()
			Expect(err).Should(Succeed())
			Expect(positions).Should(Not(BeEmpty()))
		})

		It("Dir", func() {
			positions, err := platform.Positions()

			positions = positions.Dir(true)

			Expect(err).Should(Succeed())
			Expect(positions).Should(Not(BeEmpty()))
		})

		It("Dir false", func() {
			positions, err := platform.Positions()

			positions = positions.Dir(false)

			Expect(err).Should(Succeed())
			Expect(positions).Should(BeEmpty())
		})

		It("Short", func() {
			positions, err := platform.Positions()

			positions = positions.Short()

			Expect(err).Should(Succeed())
			Expect(positions).Should(BeEmpty())
		})

		It("Long", func() {
			positions, err := platform.Positions()

			positions = positions.Long()

			Expect(err).Should(Succeed())
			Expect(positions).Should(Not(BeEmpty()))
		})

		It("Profit", func() {
			positions, err := platform.Positions()

			positions = positions.Profit()

			Expect(err).Should(Succeed())
			Expect(positions).Should(BeEmpty())
		})

		It("Lose ", func() {
			positions, err := platform.Positions()

			positions = positions.Lose()

			Expect(err).Should(Succeed())
			Expect(positions).Should(Not(BeEmpty()))
		})
	})
})
