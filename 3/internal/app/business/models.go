package business

import (
	"fmt"
)

// Allowed values for Stratum
const (
	_ Stratum = iota
	LowLow
	Low
	MediumLow
	Medium
	MediumHigh
	High
)

type Stratum uint8

func (s Stratum) Validate() error {
	switch s {
	case 0:
		return fmt.Errorf("%w: undefined stratum", ErrInvalidStratum)
	case LowLow, Low, MediumLow, Medium, MediumHigh, High:
		return nil
	}

	return fmt.Errorf("%w: %d is not a valid stratum", ErrInvalidStratum, s)
}

type Order struct {
	Products []Product
	Stratum  Stratum
}

func (o Order) Validate() (err error) {
	if o.Stratum == 0 {
		return fmt.Errorf("%w: stratum is required to process the order", ErrInvalidStratum)
	}

	if len(o.Products) == 0 {
		return fmt.Errorf("%w: order does not contain any product", ErrEmptyOrder)
	}

	for _, p := range o.Products {
		err = p.Validate()
		if err != nil {
			return err
		}
	}

	return
}

func (o Order) Summary(shippingCost float64) (summary OrderSummary) {
	for _, product := range o.Products {
		summary.Total += float64(product.Quantity) * product.Price
	}

	summary.Subtotal = summary.Total
	summary.Total += shippingCost
	summary.ShippingCost = shippingCost
	return
}

type Product struct {
	Price    float64
	Quantity uint64
}

func (p Product) Validate() (err error) {
	if p.Price == 0 && p.Quantity == 0 {
		return fmt.Errorf("%w: undefined product", ErrInvalidProduct)
	}

	return
}

type OrderSummary struct {
	Subtotal, ShippingCost, Discount, Total float64
}
