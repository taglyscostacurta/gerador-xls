package main

import (
	"github.com/labstack/echo/v4"
	"github.com/taglyscostacurta/gerador-xls/controllers"
	"github.com/taglyscostacurta/gerador-xls/db"

	//"github.com/taglyscostacurta/gerador-xls/models"
	"github.com/taglyscostacurta/gerador-xls/services"
)

func main() {

	services.Generator() // gera o arquivo XLSX

	//models.GenerateProducts()                      // gera os produtos
	db.ConnectDB()                                 // conecta com o banco de dados
	e := echo.New()                                // cria uma nova instância do Echo
	e.GET("/products", controllers.GetProducts)    // cria uma rota para a função GetProducts
	e.POST("/products", controllers.CreateProduct) // cria uma rota para a função CreateProduct

	e.Logger.Fatal(e.Start(":8000")) // inicia o servidor na porta 8000

}
