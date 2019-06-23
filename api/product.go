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

type CreateProductRequest struct {
	model.Product
	StockInit uint `json:"stockInitQty"`
}

func CreateProduct(c echo.Context) error {

	db := db.DbManager()

	request := CreateProductRequest{}
	err := JsonBodyTo(c, &request)

	if err != nil {
		return ErrorResponse(c, err)
	}

	product := model.Product(request.Product)
	stockMove := model.StockMove{
		ProductId: product.ID,
		Qty:       float32(request.StockInit),
	}

	tx := db.Begin()

	if err := tx.Create(&product).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}

	if err := tx.Create(&stockMove).Error; err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}

	if err := UpdateStockBalance(db, product.ID, float32(request.StockInit)); err != nil {
		tx.Rollback()
		return ErrorResponse(c, err)
	}

	tx.Commit()

	db.Set("gorm:auto_preload", true).First(&product, product.ID)

	return c.JSON(http.StatusCreated, product)
}
