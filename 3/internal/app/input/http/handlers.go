package http

import (
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/jikkosoft/3/internal/app/business"
	"net/http"
)

func NewHandler(processor business.OrderProcessor) (Handler, error) {
	return Handler{
		processor: processor,
	}, nil
}

type Handler struct {
	processor business.OrderProcessor
}

func (h Handler) PostOrder(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var order Order

	err = c.Bind(&order)
	if err != nil {
		return
	}

	summary, err := h.processor.ProcessOrder(ctx, order.ToBusiness())
	if err != nil {
		return
	}

	return c.JSON(http.StatusCreated, NewSummary(summary))
}
