package database

import (
	"log"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {
	conexao := "your database"
	DB, err = gorm.Open(oracle.Open(conexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
