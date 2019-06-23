package api

import (
	"ezsale/model"

	"github.com/labstack/echo"
)

func GetSellingUnit(c echo.Context) error {
	return SearchMaster(c, &[]model.SellingUnit{})
}

func CreateSellingUnit(c echo.Context) error {
	return CreateMasterData(c, &model.SellingUnit{})
}

func UpdateSellingUnit(c echo.Context) error {
	return UpdateMasterData(c, &model.SellingUnit{})
}

func DeleteSellingUnit(c echo.Context) error {
	return DeleteMaster(c, &model.SellingUnit{})
}
