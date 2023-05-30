package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Connect() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=MSpg_473 dbname=bookstore_db port=5432 sslmode=disable TimeZone=Asia/Almaty"),
		&gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDB() *gorm.DB {
	return database
}
