package main

import (
	"github.com/labstack/echo/v4"
	"github.com/taglyscostacurta/gerador-xls/db"
	"github.com/taglyscostacurta/gerador-xls/services"
)

func main() {

	services.Generator() // gera o arquivo XLSX

	db.GenerateProducts()                 // gera os produtos
	e := echo.New()                       // cria uma nova instância do Echo
	e.GET("/products", db.GetProducts)    // cria uma rota para a função GetProducts
	e.POST("/products", db.CreateProduct) // cria uma rota para a função CreateProduct
	e.Logger.Fatal(e.Start(":8000"))      // inicia o servidor na porta 8000

}
