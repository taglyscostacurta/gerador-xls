package controllers

import (
	"database/sql" // importa o pacote de banco de dados

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // importa o driver do sqlite3
	"github.com/taglyscostacurta/gerador-xls/models"
)

func GetProducts(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return err
		}
		products = append(products, product)
	}
	return c.JSON(200, products)
}

func CreateProduct(c echo.Context) error {
	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		return err
	}
	defer db.Close()

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO products(name, price) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {

		return err
	}

	product.ID = uint(id)
	return c.JSON(201, product)
}
