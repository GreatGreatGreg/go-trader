package platform

import "github.com/santacruz123/go-trader/trades"

// Platformer - interface of platform
type Platformer interface {

	// Symbol - creates new Symbol
	Symbol(string) (*trades.Symbol, error)

	// Orders - gets all orders
	Orders() (trades.Orders, error)

	// Positions - gets all positions
	Positions() (trades.Positions, error)

	// Order - sends order
	Order(trades.Order) (uint, error)

	// Cancel - cancels open order
	Cancel(uint) error

	// CancelAll - cancels all open order
	CancelAll() error

	// Modify - modifies open order
	Modify(uint, trades.Order) error
}
