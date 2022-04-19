package handler

import (
	"fmt"
	"net/http"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type tag struct {
	service service.TagService
}

func NewTagHandler(e *echo.Echo, s service.TagService) {
	handler := &tag{
		service: s,
	}

	e.POST("/api/tags", handler.Create, auth())
	e.GET("/api/tags", handler.Read, auth())
	e.PUT("/api/tags/:id", handler.Update, auth())
	e.GET("/api/tags/:id", handler.ReadByID, auth())
	e.DELETE("/api/tags/:id", handler.Delete, auth())
}

func (p *tag) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
	)

	data, err := p.service.Read(ctx, paging(c))
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.TagEntity),
		Data:    data,
	})
}

func (p *tag) ReadByID(c echo.Context) error {
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
		Message: fmt.Sprintf(constans.MessageSuccessReadByID, constans.TagEntity, id),
		Data:    data,
	})
}

func (p *tag) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req model.Tag
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
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.TagEntity),
		Data:    nil,
	})
}

func (p *tag) Update(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
		req model.Tag
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
		Message: fmt.Sprintf(constans.MessageSuccessUpdate, constans.TagEntity, id),
		Data:    response,
	})
}

func (p *tag) Delete(c echo.Context) error {
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
		Message: fmt.Sprintf(constans.MessageSuccessDelete, constans.TagEntity, id),
		Data:    nil,
	})
}
