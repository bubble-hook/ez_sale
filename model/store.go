package model

type Store struct {
	Model
	StoreID     string `gorm:"unique;not null"json:"storeId"`
	OwnerUserID uint   `json:"ownerUserID"`
	Name        string `json:"name"`
}

type StoreUser struct {
	ID      uint `gorm:"primary_key"json:"id"`
	StoreID uint
	UserID  uint
}
