package handler

import (
	"strconv"

	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/labstack/echo/v4"
)

func paging(c echo.Context) model.Paging {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	return model.Paging{
		Limit:  int64(limit),
		Offset: int64(offset),
	}
}
