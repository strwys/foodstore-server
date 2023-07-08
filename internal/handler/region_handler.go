package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cecepsprd/foodstore-server/constans"
	"github.com/cecepsprd/foodstore-server/internal/model"
	"github.com/cecepsprd/foodstore-server/utils"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
)

type region struct {
}

func NewRegionHandler(e *echo.Echo) {
	handler := &region{}

	e.GET("/api/region/provinsi", handler.ReadAllProvince, auth())
	e.GET("/api/region/kabupaten", handler.ReadAllRegencies, auth())
	e.GET("/api/region/kecamatan", handler.ReadAllDistricts, auth())
	e.GET("/api/region/desa", handler.ReadAllVillages, auth())
}

func (p *region) ReadAllProvince(c echo.Context) error {
	var provinces []model.Province

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
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, `Province`),
		Data:    provinces,
	})
}

func (p *region) ReadAllRegencies(c echo.Context) error {
	var regencies []model.Regency
	var kodeInduk = c.QueryParam("kode_induk")

	f, err := os.Open("./datastore/data-wilayah/regencies.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &regencies); err != nil {
		fmt.Println(err)
		return err
	}

	var filteredRegency []model.Regency
	for _, val := range regencies {
		if val.ProvinceCode == kodeInduk {
			filteredRegency = append(filteredRegency, val)
		}
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, `Regency`),
		Data:    filteredRegency,
	})
}

func (p *region) ReadAllDistricts(c echo.Context) error {
	var districts []model.District
	var kodeInduk = c.QueryParam("kode_induk")

	f, err := os.Open("./datastore/data-wilayah/districts.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &districts); err != nil {
		fmt.Println(err)
		return err
	}

	var filteredDistrict []model.District
	for _, val := range districts {
		if val.RegencyCode == kodeInduk {
			filteredDistrict = append(filteredDistrict, val)
		}
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, `District`),
		Data:    filteredDistrict,
	})
}

func (p *region) ReadAllVillages(c echo.Context) error {
	var villages []model.Village
	var kodeInduk = c.QueryParam("kode_induk")

	f, err := os.Open("./datastore/data-wilayah/villages.csv")
	if err != nil {
		return c.JSON(utils.SetHTTPStatusCode(err), model.ResponseError{Message: err.Error()})
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &villages); err != nil {
		fmt.Println(err)
		return err
	}

	var filteredVillages []model.Village
	for _, val := range villages {
		if val.DistrictCode == kodeInduk {
			filteredVillages = append(filteredVillages, val)
		}
	}

	return c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf(constans.MessageSuccessReadAll, `Village`),
		Data:    filteredVillages,
	})
}
