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

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(e *echo.Echo, us service.UserService) {
	handler := &UserHandler{
		userService: us,
	}

	e.POST("/api/users", handler.Create)
	e.GET("/api/users", handler.ReadAllUser, auth())
}

func (u *UserHandler) ReadAllUser(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := u.userService.Read(ctx)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.UserEntity),
		Data:    data,
	})
}

func (u *UserHandler) Create(c echo.Context) error {
	var req model.User
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	err = u.userService.Create(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.UserEntity),
		Data:    nil,
	})
}
