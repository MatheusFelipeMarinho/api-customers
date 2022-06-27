package database

import (
	"api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Address{})
}
