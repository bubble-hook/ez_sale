package api

import (
	"ezsale/db"
	"ezsale/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetProduct(c echo.Context) error {
	return SearchMaster(c, model.Product{})
}

func CreateProduct(c echo.Context) error {

	rq := model.ProductCreateRequest{}

	err := JsonBodyTo(c, &rq)
	db := db.DbManager()
	if err != nil {
		return ErrorResponse(c, err)
	}

	product := model.Product(rq.Product)
	goods := model.Goods{}
	goods.MasterData = model.MasterData(product.MasterData)
	goods.UnitPrice = rq.UnitPrice
	goods.UnitQuantityID = rq.MainUnitQuantityID

	tx := db.Begin()

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}

	goods.ProductID = product.ID

	if err := tx.Create(&goods).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}

	tx.Commit()

	db.Set("gorm:auto_preload", true).First(&product, product.ID)

	return c.JSON(http.StatusCreated, product)
}
