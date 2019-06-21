package api

import (
	"ezsale/model"

	"github.com/labstack/echo"
)

func GetUnitQuantity(c echo.Context) error {
	return SearchMaster(c, &[]model.UnitQuantity{})
}

func CreateUnitQuantity(c echo.Context) error {
	return CreateMasterData(c, &model.UnitQuantity{})
}

func UpdateUnitQuantity(c echo.Context) error {
	return UpdateMasterData(c, &model.UnitQuantity{})
}

func DeleteUnitQuantity(c echo.Context) error {
	return DeleteMaster(c, &model.UnitQuantity{})
}
