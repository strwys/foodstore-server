package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/internal/service"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/cecepsprd/foodstore-server/utils/logger"
	"golang.org/x/sync/errgroup"

	"github.com/labstack/echo/v4"
)

type product struct {
	service service.ProductService
}

func NewProductHandler(e *echo.Echo, s service.ProductService) {
	handler := &product{
		service: s,
	}

	e.POST("/api/products", handler.Create, auth())
	e.GET("/api/products", handler.Read)
	e.PUT("/api/products/:id", handler.Update, auth())
	e.GET("/api/products/:id", handler.ReadByID, auth())
	e.DELETE("/api/products/:id", handler.Delete, auth())
}

func (p *product) Read(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = model.ReadProductRequest{}
	)

	err := utils.DecodeQueryParams(c.QueryString(), "params", &req)
	if err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	data, total, err := p.service.Read(ctx, req)
	if err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.ReadAllProductResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.ProductEntity),
		Data:    data,
		Total:   total,
	})
}

func (p *product) ReadByID(c echo.Context) error {
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
		Message: fmt.Sprintf(constans.MessageSuccessReadByID, constans.ProductEntity, id),
		Data:    data,
	})
}

func (p *product) Create(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req model.ProductRequest
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	req.ImageURL, err = p.uploadImage(c, "")
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	err = p.service.Create(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, model.APIResponse{
		Code:    http.StatusCreated,
		Message: fmt.Sprintf(constans.MessageSuccessCreate, constans.ProductEntity),
		Data:    nil,
	})
}

func (p *product) Update(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = c.Param("id")
		req model.ProductRequest
	)

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	req.ImageURL, err = p.uploadImage(c, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	response, err := p.service.Update(ctx, req)
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessUpdate, constans.ProductEntity, id),
		Data:    response,
	})
}

func (p *product) Delete(c echo.Context) error {
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
		Message: fmt.Sprintf(constans.MessageSuccessDelete, constans.ProductEntity, id),
		Data:    nil,
	})
}

func (p *product) uploadImage(c echo.Context, productID string) (imageURL string, err error) {
	g, ctx := errgroup.WithContext(c.Request().Context())

	// upload new image
	g.Go(func() error {
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}

		imageFormat := strings.Split(file.Filename, ".")
		imageURL = fmt.Sprintf(constans.BaseImagePath, time.Now().UnixNano(), imageFormat[1])
		dst, err := os.Create(imageURL)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return nil
	})

	// remove old image
	g.Go(func() error {
		prd, _ := p.service.ReadByID(ctx, productID)
		if !reflect.ValueOf(prd).IsNil() {
			if !reflect.ValueOf(prd.ImageURL).IsZero() {
				os.Remove(prd.ImageURL)
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return "", err
	}

	return imageURL, nil
}
