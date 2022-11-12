package handler

import (
	"fmt"
	"net/http"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type address struct {
	service service.AddressService
}

func NewDeliveryAddressHandler(e *echo.Echo, s service.AddressService) {
	handler := &address{
		service: s,
	}

	e.POST("/api/delivery-addresses", handler.Create, auth())
	e.GET("/api/delivery-addresses", handler.Read, auth())
	e.PUT("/api/delivery-addresses/:id", handler.Update, auth())
	e.DELETE("/api/delivery-addresses/:id", handler.Delete, auth())
}

func (p *address) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	data, err := p.service.Read(ctx)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.DeliveryAddressEntity),
		Data:    data,
	})
}

func (p *address) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req model.DeliveryAddress
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	err = p.service.Create(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.APIResponse{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.DeliveryAddressEntity),
		Data:    nil,
	})
}

func (p *address) Update(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
		req model.DeliveryAddress
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	req.ID, _ = primitive.ObjectIDFromHex(id)
	response, err := p.service.Update(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessUpdate, constans.DeliveryAddressEntity, id),
		Data:    response,
	})
}

func (p *address) Delete(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
	)

	err := p.service.Delete(ctx, id)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessDelete, constans.DeliveryAddressEntity, id),
		Data:    nil,
	})
}
