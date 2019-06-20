package model

import "time"

type Model struct {
	ID          uint       `gorm:"primary_key"json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
	CreatedById int        `json:"createdById"`
	UpdatedById int        `json:"updatedById"`
}

type MasterData struct {
	Name string `json:"name"`
	Code string `gorm:"unique;not null"json:"code"`
}
