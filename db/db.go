package db

import (
	"log"

	"github.com/taglyscostacurta/gerador-xls/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	DB, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.Product{}) // cria a tabela products no banco de dados
}
