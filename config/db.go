package config

import (
	"golangAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		// 也可以 import log 之後用log.Fatal(err)替代
		panic(err)

	}
	db.AutoMigrate(&models.User{})
	DB = db
}