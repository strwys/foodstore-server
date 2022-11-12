package handler

import (
	"fmt"
	"net/http"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
	"github.com/labstack/echo"
)

type order struct {
	service service.OrderService
}

func NewOrderHandler(e *echo.Echo, s service.OrderService) {
	handler := &order{
		service: s,
	}

	e.POST("/api/orders", handler.Create, auth())
	e.GET("/api/orders", handler.Read, auth())
}

func (h *order) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req model.CreateOrderRequest
	)

	err := c.Bind(&req)
	if err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	req.User = utils.GetUserByContext(c)

	order, err := h.service.Create(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.APIResponse{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.OrderEntity),
		Data:    order,
	})
}

func (h *order) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	orders, err := h.service.Read(ctx, utils.GetUserIDByContext(c))
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.OrderEntity),
		Data:    orders,
	})
}
