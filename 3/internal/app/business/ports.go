package business

import (
	"context"
)

// Ports for drive adapters
type (
	// OrderProcessor defines a way to process and order to get an OrderSummary
	OrderProcessor interface {
		ProcessOrder(context.Context, Order) (OrderSummary, error)
	}
)

// Ports for driven adapters
type (
	// ShippingCostCalculator defines a way to calculate the shipping cost for an Order based on the Stratum
	ShippingCostCalculator interface {
		CalculateShippingCost(context.Context, Stratum) (float64, error)
	}
)
