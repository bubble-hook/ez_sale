package api

import (
	"ezsale/model"

	"github.com/labstack/echo"
)

func GetProductCategory(c echo.Context) error {
	return SearchMaster(c, &[]model.ProductCategory{})
}

func CreateProductCategory(c echo.Context) error {
	return CreateMasterData(c, &model.ProductCategory{})
}

func UpdateProductCategory(c echo.Context) error {
	return UpdateMasterData(c, &model.ProductCategory{})
}

func DeleteProductCategory(c echo.Context) error {
	return DeleteMaster(c, &model.ProductCategory{})
}
