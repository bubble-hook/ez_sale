package api

import (
	"ezsale/db"
	"ezsale/model"
	"net/http"

	"github.com/labstack/echo"
)

type ProcessOrderRequest struct {
	Discount float32                   `json:"discount" `
	Items    []ProcessOrderRequestItem `json:"items" validate:"required"`
}
type ProcessOrderRequestItem struct {
	ProductId    uint    `json:"productId" validate:"required"`
	SellingPrice float32 `json:"sellingPrice" validate:"required"`
	Qty          uint    `json:"qty" validate:"required"`
}

func ProcessOrder(c echo.Context) error {

	db := db.DbManager()
	request := ProcessOrderRequest{}
	err := JsonBodyTo(c, &request)
	if err != nil {
		return ErrorResponse(c, err)
	}

	order := model.Order{}
	orderItems := []model.OrderDetail{}

	for _, item := range request.Items {

		product := model.Product{}

		if err := db.First(&product, item.ProductId).Error; err != nil {
			return ErrorResponse(c, err)
		}

		orderDetail := model.OrderDetail{
			ProductName:  product.Name,
			ProductId:    product.ID,
			Qty:          item.Qty,
			SellingPrice: item.SellingPrice,
		}

		orderItems = append(orderItems, orderDetail)
		order.Amount += orderDetail.SellingPrice * float32(orderDetail.Qty)
	}

	if order.Discount > 0 && order.Amount > 0 {
		order.DiscountVal = order.Amount * order.Discount / 100
	}

	order.NetAmount = order.Amount - order.DiscountVal
	order.ItemsCount = uint(len(orderItems))
	order.OrderItems = orderItems

	return c.JSON(http.StatusOK, order)
}
