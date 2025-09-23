package business_test

import (
	"context"
	. "github.com/yael-castro/jikkosoft/3/internal/app/business"
	"github.com/yael-castro/jikkosoft/3/internal/app/output/grpc"
	"reflect"
	"strconv"
	"testing"
)

func TestOrderProcessor_ProcessOrder(t *testing.T) {
	cases := [...]struct {
		ctx      context.Context
		order    Order
		expected OrderSummary
	}{
		{
			order: Order{
				Stratum: High,
				Products: []Product{
					{Quantity: 1, Price: 100},
					{Quantity: 2, Price: 200},
				},
			},
			expected: OrderSummary{
				Subtotal:     500,
				ShippingCost: 10,
				Total:        510,
			},
		},
	}

	processor, err := NewOrderProcessor(grpc.NewShippingCalculator())
	if err != nil {
		t.Fatal(err)
	}

	for i, testCase := range cases {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			summary, err := processor.ProcessOrder(testCase.ctx, testCase.order)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(testCase.expected, summary) {
				t.Fatalf("order summary %+v, want %+v", summary, testCase.order)
			}
		})
	}
}
