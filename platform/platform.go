package platform

import "../trades"

// Platform - interface of platform
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

	// Modify - modifies open order
	Modify(uint, trades.Order) error
}
