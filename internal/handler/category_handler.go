package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/strwys/foodstore-server/constans"
	"github.com/strwys/foodstore-server/internal/model"
	"github.com/strwys/foodstore-server/internal/service"
	"github.com/strwys/foodstore-server/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type category struct {
	service service.CategoryService
}

func NewCategoryHandler(e *echo.Echo, s service.CategoryService) {
	handler := &category{
		service: s,
	}

	e.POST("/api/category", handler.Create, auth())
	e.GET("/api/category", handler.Read, auth())
	e.PUT("/api/category/:id", handler.Update, auth())
	e.GET("/api/category/:id", handler.ReadByID, auth())
	e.DELETE("/api/category/:id", handler.Delete, auth())
}

func (p *category) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	data, err := p.service.Read(ctx, paging(c))
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.CategoryEntity),
		Data:    data,
	})
}

func (p *category) ReadByID(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
	)

	data, err := p.service.ReadByID(ctx, id)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadByID, constans.CategoryEntity, id),
		Data:    data,
	})
}

func (p *category) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req model.Category
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
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.CategoryEntity),
		Data:    nil,
	})
}

func (p *category) Update(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
		req model.Category
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
		Message: fmt.Sprintf(constans.MessageSuccessUpdate, constans.CategoryEntity, id),
		Data:    response,
	})
}

func (p *category) Delete(c echo.Context) error {
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
		Message: fmt.Sprintf(constans.MessageSuccessDelete, constans.CategoryEntity, id),
		Data:    nil,
	})
}
