package models

type Product struct {
	ID    uint    `json:"id" gorm:"primary_key"` // id do produto
	Name  string  `json:"name"`                  // nome do produto
	Price float64 `json:"price"`                 // preço do produto

}

var Products []Product // slice de produtos
