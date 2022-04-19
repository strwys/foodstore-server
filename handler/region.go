package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/model"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo"
)

type region struct {
}

func NewRegionHandler(e *echo.Echo) {
	handler := &region{}

	e.GET("/api/region/provinces", handler.ReadAllProvince)
	e.GET("/api/region/districts", handler.ReadAllDistricts)
	e.GET("/api/region/regencies", handler.ReadAllRegencies)
	e.GET("/api/region/villages", handler.ReadAllVillages)
}

func (p *region) ReadAllProvince(c echo.Context) error {
	var provinces []*model.Province

	f, err := os.Open("./datastore/data-wilayah/provinces.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &provinces); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.TagEntity),
		Data:    provinces,
	})
}

func (p *region) ReadAllRegencies(c echo.Context) error {
	var regencies []*model.Regency

	f, err := os.Open("./datastore/data-wilayah/regencies.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &regencies); err != nil {
		fmt.Println(err)
		return err
	}

	var filteredRegency []*model.Regency
	for _, val := range regencies {
		if val.ProvinceCode == c.Param("province_code") {
			filteredRegency = append(filteredRegency, val)
		}
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.TagEntity),
		Data:    filteredRegency,
	})
}

func (p *region) ReadAllDistricts(c echo.Context) error {
	var districts []*model.District

	f, err := os.Open("./datastore/data-wilayah/districts.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &districts); err != nil {
		fmt.Println(err)
		return err
	}

	var filteredDistrict []*model.District
	for _, val := range districts {
		if val.RegencyCode == c.Param("regency_code") {
			filteredDistrict = append(filteredDistrict, val)
		}
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.TagEntity),
		Data:    filteredDistrict,
	})
}

func (p *region) ReadAllVillages(c echo.Context) error {
	var villages []*model.Village

	f, err := os.Open("./datastore/data-wilayah/villages.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &villages); err != nil {
		fmt.Println(err)
		return err
	}

	var filteredVillages []*model.Village
	for _, val := range villages {
		if val.DistrictCode == c.Param("district_code") {
			filteredVillages = append(filteredVillages, val)
		}
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, constans.TagEntity),
		Data:    filteredVillages,
	})
}
