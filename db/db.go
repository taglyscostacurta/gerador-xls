package db

type Product struct {
	Name  string
	Price float64
}

var products []Product

func CreateProduct(product Product) {
	products = append(products, product)
}
