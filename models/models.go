package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primarykey" AutoIncrement:"true"`
	Name      string         `json:"name" validate:"nonzero"`
	Price     uint           `json:"price" validate:"nonzero"`
	Code      string         `json:"code" validate:"nonzero"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

var Products []Product // slice de produtos
