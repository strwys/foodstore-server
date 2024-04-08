package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/strwys/foodstore-server/internal/model"
)

func paging(c echo.Context) model.Paging {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	return model.Paging{
		Limit:  int64(limit),
		Offset: int64(offset),
	}
}
