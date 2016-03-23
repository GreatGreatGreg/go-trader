package bitfinex_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBitfinex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bitfinex Suite")
}
