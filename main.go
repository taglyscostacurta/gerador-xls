package main

import (
	"github.com/labstack/echo/v4"
	"github.com/taglyscostacurta/gerador-xls/db"
	_ "github.com/taglyscostacurta/gerador-xls/services"
)

func main() {

	//services.Generator()

	db.CreateProducts()
	e := echo.New()
	e.GET("/products", db.GetProducts)
	e.Logger.Fatal(e.Start(":8000"))

}
