package database

import (
	"log"

	"github.com/jonathantx/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	connection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
