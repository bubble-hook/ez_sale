package api

import (
	"ezsale/db"
	"ezsale/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetProductCategory(c echo.Context) error {
	return SearchMaster(c, &[]model.ProductCategory{})
}

func CreateProductCategory(c echo.Context) error {
	db := db.DbManager()
	m := model.ProductCategory{}
	err := JsonBodyTo(c, &m)
	if err != nil {
		return ErrorResponse(c, err)
	}
	count := 0
	db.Model(&model.ProductCategory{}).Where("code = ?", m.Code).Count(&count)
	if count > 0 {
		return ErrorResponseMessage(c, http.StatusBadRequest, "Duplicated Code")
	}
	db.Create(&m)
	db.Save(&m)
	return c.JSON(http.StatusCreated, m)
}
