package grpc

import (
	"context"
	"fmt"
	"github.com/yael-castro/jikkosoft/3/internal/app/business"
)

func NewShippingCalculator() business.ShippingCostCalculator {
	return shippingCostCalculator{}
}

type shippingCostCalculator struct{}

func (s shippingCostCalculator) CalculateShippingCost(ctx context.Context, stratum business.Stratum) (float64, error) {
	// Here I am simulating query some external API
	switch stratum {
	case business.MediumHigh, business.High:
		return 10, nil
	case business.MediumLow, business.Medium:
		return 20, nil
	case business.LowLow, business.Low:
		return 30, nil
	}

	return 0, fmt.Errorf("%w: %v is an invalid stratum", business.ErrInvalidStratum, stratum)
}
