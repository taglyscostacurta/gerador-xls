package db

import (
	"github.com/labstack/echo/v4"
)

type Product struct {
	Name  string
	Price float64
}

var products []Product

func CreateProducts() {
	products = append(products, Product{Name: "Product 1", Price: 10.00})
	products = append(products, Product{Name: "Product 2", Price: 20.00})
	products = append(products, Product{Name: "Product 3", Price: 30.00})
}

func GetProducts(c echo.Context) error {
	return c.JSON(200, products)
}
