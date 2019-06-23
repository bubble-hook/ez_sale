package model

type Order struct {
	Model
	OrderId     string        `gorm:"unique;not null"json:"OrderId"`
	Discount    float32       `json"discount"`
	DiscountVal float32       `json:"discountVal"`
	ItemsCount  uint          `json:"itemsCount"`
	Amount      float32       `json:"amount"`
	NetAmount   float32       `json:"netAmount"`
	OrderItems  []OrderDetail `gorm:"foreignkey:OrderId" json:"orderItems"`
}

type OrderDetail struct {
	ID           uint    `gorm:"primary_key"json:"id"`
	OrderId      uint    `orderId`
	ProductId    uint    `json:"productId"`
	ProductName  string  `json:"productName"`
	SellingPrice float32 `json"sellingPrice"`
	Qty          uint    `json:"qty"`
}
