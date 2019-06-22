package model

import "time"

type UnitQuantity struct {
	Model
	MasterData
	Qty float32 `json:"qty"`
}

type ProductCategory struct {
	Model
	MasterData
}

type ProductCreateRequest struct {
	Product
	UnitPrice float32 `json:"unitPrice"`
}

type Product struct {
	Model
	MasterData
	ProductCategoryID  uint            `json:"productCategotyID"`
	ProductCategory    ProductCategory `gorm:"foreignkey:ProductCategoryID"json:"productCategory"`
	MainUnitQuantityID uint            `json:"mainUnitQuantityID"`
	UnitQuantity       UnitQuantity    `gorm:"foreignkey:MainUnitQuantityID"json:"mainUnitQuantity"`
	GoodsItems         []Goods         `gorm:"foreignkey:ProductID"json:"goodsItems"`
}

type Goods struct {
	Model
	MasterData
	ProductID uint `json:"productId"`
	//Product        Product      `gorm:"foreignkey:ProductCategoryID"json:"product"`
	UnitQuantityID uint         `json:"unitQuantityID"`
	UnitQuantity   UnitQuantity `gorm:"foreignkey:unitQuantityID"json:"unitQuantity"`
	UnitPrice      float32      `json:"unitPrice"`
}

type StockMove struct {
	ID        uint      `gorm:"primary_key"json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductId uint      `json:"productId"`
	Qty       float32   `json:"qty"`
}

type StockBalance struct {
	ID         uint      `gorm:"primary_key"json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ProductId  uint      `json:"productId"`
	PrvQty     float32
	CurrentQty float32
}
