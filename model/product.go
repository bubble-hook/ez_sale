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

type Product struct {
	Model
	MasterData
	ProductCategoryID  uint            `json:"productCategotyID"`
	ProductCategory    ProductCategory `gorm:"foreignkey:ProductCategoryID"json:"productCategory"`
	MainUnitQuantityID uint            `json:"mainUnitQuantityID"`
	UnitQuantity       UnitQuantity    `gorm:"foreignkey:MainUnitQuantityID"json:"mainUnitQuantity"`
}

type Goods struct {
	Model
	MasterData
	ProductID      uint         `json:"productId"`
	Product        Product      `gorm:"foreignkey:ProductCategoryID"json:"product"`
	UnitQuantityID uint         `json:"unitQuantityID"`
	UnitQuantity   UnitQuantity `gorm:"foreignkey:"ProductCategoryID"json:"unitQuantity"`
	UnitPrice      float32      `json:"unitPrice"`
}

type StockMove struct {
	ID        uint      `gorm:"primary_key"json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProductId uint      `json:"productId"`
	Qty       float32   `json:"qty"`
}