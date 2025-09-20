package business

import (
	"context"
	"errors"
)

func NewOrderProcessor(calculator ShippingCostCalculator) (OrderProcessor, error) {
	if calculator == nil {
		return nil, errors.New("missing calculator")
	}

	return orderProcessor{
		calculator: calculator,
	}, nil
}

type orderProcessor struct {
	calculator ShippingCostCalculator
}

func (o orderProcessor) ProcessOrder(ctx context.Context, order Order) (summary OrderSummary, err error) {
	err = order.Validate()
	if err != nil {
		return
	}

	// Calculating shipping cost using an external API or something like that
	shippingCost, err := o.calculator.CalculateShippingCost(ctx, order.Stratum)
	if err != nil {
		return
	}

	summary = order.Summary(shippingCost)
	return
}
