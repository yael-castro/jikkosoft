package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoutes(e *echo.Echo, handler Handler, checks ...func(context.Context) error) {
	g := e.Group("/v1/orders")

	// Setting health check
	e.GET("/v1/health", health(checks...))

	// Setting user routes
	g.POST("", handler.PostOrder)
}

func health(checks ...func(context.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		for _, check := range checks {
			err := check(ctx)
			if err != nil {
				return c.JSON(http.StatusServiceUnavailable, struct{}{})
			}
		}

		return c.JSON(http.StatusOK, struct{}{})
	}
}
