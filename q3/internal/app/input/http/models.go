package http

import "github.com/yael-castro/jikkosoft/3/internal/app/business"

type Order struct {
	Products []business.Product `json:"products"`
	Stratum  business.Stratum   `json:"stratum"`
}

// ToBusiness Translate Order into an object that the business layer can understand (business.Order).
func (o Order) ToBusiness() business.Order {
	return business.Order(o)
}

// NewSummary translating business.OrderSummary into OrderSummary ...
func NewSummary(summary business.OrderSummary) OrderSummary {
	return OrderSummary(summary)
}

type OrderSummary struct {
	Subtotal     float64 `json:"subtotal"`
	ShippingCost float64 `json:"shipping_cost"`
	Discount     float64 `json:"discount"`
	Total        float64 `json:"total"`
}
