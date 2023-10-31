package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	stringConnection := "host=localhost user=caiopcabral password=0102 dbname=api port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Error while connecting to Database!")
	}
}
