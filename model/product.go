package model

import "time"

type SellingUnit struct {
	Model
	MasterData
}

type ProductCategory struct {
	Model
	MasterData
}

type Product struct {
	Model
	MasterData
	ProductCategoryID uint            `json:"productCategotyID"`
	ProductCategory   ProductCategory `gorm:"foreignkey:ProductCategoryID"json:"productCategory"`
	SellingUnitID     uint            `json:"sellingUnitID"`
	SellingUnit       SellingUnit     `gorm:"foreignkey:SellingUnitID"json:"sellingUnitID"`
	SellingPrice      float32         `json:"sellingPrice"`
	Cost              float32         `json:"cost"`
}

type StockMove struct {
	ID        uint      `gorm:"primary_key"json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Refer     uint
	ProductId uint    `json:"productId"`
	Qty       float32 `json:"qty"`
}

type StockBalance struct {
	ID         uint      `gorm:"primary_key"json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ProductId  uint      `json:"productId"`
	PrvQty     float32
	CurrentQty float32
}
