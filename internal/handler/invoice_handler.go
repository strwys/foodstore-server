package handler

import (
	"fmt"
	"net/http"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/labstack/echo/v4"
)

type invoice struct {
	service service.InvoiceService
}

func NewInvoiceHandler(e *echo.Echo, s service.InvoiceService) {
	handler := &invoice{
		service: s,
	}

	e.GET("/api/invoice/:order_id", handler.Read, auth())
}

func (h *invoice) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	orderid := c.Param("order_id")
	invoice, err := h.service.Read(ctx, orderid)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.APIResponse{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.InvoiceEntity),
		Data:    invoice,
	})
}
