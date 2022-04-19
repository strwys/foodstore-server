package handler

import (
	"fmt"
	"net/http"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(e *echo.Echo, authService service.AuthService, userService service.UserService) {
	handler := &AuthHandler{
		authService: authService,
		userService: userService,
	}

	e.POST("/api/auth/login", handler.Login)
	e.POST("/api/auth/register", handler.Register)
}

func (ah *AuthHandler) Login(c echo.Context) error {
	var (
		req = model.LoginRequest{}
		ctx = c.Request().Context()
	)

	err := c.Bind(&req)
	if err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	fmt.Println("->", req)

	response, err := ah.authService.Login(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *AuthHandler) Register(c echo.Context) error {
	var (
		ctx  = c.Request().Context()
		req  = model.RegisterRequest{}
		user = model.User{}
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

	if err = utils.MappingRequest(req, &user); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	err = u.userService.Create(ctx, user)
	if err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.UserEntity),
		Data:    nil,
	})
}

func auth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString("APP_JWT_SECRET")),
	})
}

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		roles := claims["roles"].(string)
		if roles != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
