package container

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yael-castro/jikkosoft/3/internal/app/business"
	"github.com/yael-castro/jikkosoft/3/internal/app/input/http"
	"github.com/yael-castro/jikkosoft/3/internal/app/output/grpc"
)

func New() *Container {
	return new(Container)
}

type Container struct{}

func (c Container) Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case **echo.Echo:
		return c.injectEcho(ctx, a)
	}

	return fmt.Errorf("unsupported type: %T", a)
}

func (c Container) injectEcho(_ context.Context, e **echo.Echo) (err error) {
	// Secondary adapters
	calculator := grpc.NewShippingCalculator()

	// Business logic
	processor, err := business.NewOrderProcessor(calculator)
	if err != nil {
		return err
	}

	// Primary adapters
	handlers, err := http.NewHandler(processor)
	if err != nil {
		return err
	}

	// Building echo.Echo
	n := echo.New()

	// Setting error handler
	n.HTTPErrorHandler = http.ErrorHandler(n.HTTPErrorHandler)

	// Setting middlewares
	n.Use(middleware.Recover(), middleware.Logger())

	// Setting health checks
	dbCheck := func(ctx context.Context) error {
		return nil // Here I am simulating a ping to some DB or API
	}

	// Setting http routes
	http.SetRoutes(n, handlers, dbCheck)

	// Disabling initial logs
	n.HideBanner = true
	n.HidePort = true

	*e = n
	return
}
