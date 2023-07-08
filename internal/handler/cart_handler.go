package handler

import (
	"fmt"
	"net/http"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
	"github.com/labstack/echo/v4"
)

type cart struct {
	service service.CartService
}

func NewCartHandler(e *echo.Echo, s service.CartService) {
	handler := &cart{
		service: s,
	}

	e.GET("/api/carts", handler.Read, auth())
	e.PUT("/api/carts", handler.Update, auth())
	e.DELETE("/api/carts/:item_id", handler.DeleteByItemID, auth())
}

func (h *cart) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	carts, err := h.service.Read(ctx, utils.GetUserIDByContext(c))
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.CartEntity),
		Data:    carts,
	})
}

func (h *cart) Update(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		req  = model.UpdateCartItemRequest{}
		cart = []model.CartItem{}
	)

	if err := c.Bind(&req); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	if err := utils.MappingRequest(req.Items, &cart); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.ResponseError{Message: err.Error()})
	}

	userid := utils.GetUserIDByContext(c)

	if err := h.service.Update(ctx, userid, cart); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessUpdateCart),
		Data:    nil,
	})
}

func (h *cart) DeleteByItemID(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	itemID := utils.ConvertPrimitiveID(c.Param("item_id"))

	err := h.service.DeleteByItemID(ctx, utils.GetUserIDByContext(c), itemID)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessDelete, constans.CartEntity, itemID),
		Data:    nil,
	})
}
