package db

import (
	"database/sql" // importa o pacote de banco de dados

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // importa o driver do sqlite3
)

type Product struct {
	Name  string  `json:"name"`  // nome do produto
	Price float64 `json:"price"` // preço do produto

}

var products []Product // slice de produtos

func GenerateProducts() { // função que gera os produtos
	products = append(products, Product{Name: "Product 1", Price: 10.00})
	products = append(products, Product{Name: "Product 2", Price: 20.00})
	products = append(products, Product{Name: "Product 3", Price: 30.00})
}

func GetProducts(c echo.Context) error { // função que retorna os produtos
	return c.JSON(200, products) // retorna os produtos em formato JSON
}

func CreateProduct(c echo.Context) error { // função que cria um novo produto
	product := new(Product)                 // cria um novo produto
	if err := c.Bind(product); err != nil { // se ocorrer algum erro, retorna o erro
		return err
	}
	products = append(products, *product) // adiciona o produto no slice de produtos
	saveProducts(*product)                // salva o produto no banco de dados
	return c.JSON(200, product)           // retorna o produto em formato JSON
}

func saveProducts(product Product) error {
	db, err := sql.Open("sqlite3", "products.db") // aqui abre a conexão com o banco de dados
	if err != nil {
		return err // caso ocorra algum erro, retorna o erro
	}
	defer db.Close() // fecha a conexão com o banco de dados

	stmt, err := db.Prepare("INSERT INTO products(name, price) values($1, $2)") // prepara a query
	if err != nil {
		return err // caso ocorra algum erro, retorna o erro
	}

	_, err = stmt.Exec(product.Name, product.Price) // executa a query
	if err != nil {
		return err // caso ocorra algum erro, retorna o erro
	}
	return nil // caso não ocorra nenhum erro, retorna nil

}
