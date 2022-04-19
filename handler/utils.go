package handler

import (
	"strconv"

	"github.com/cecepsprd/foodstore-server/model"
	"github.com/labstack/echo"
)

func paging(c echo.Context) model.Paging {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	return model.Paging{
		Limit:  int64(limit),
		Offset: int64(offset),
	}
}
