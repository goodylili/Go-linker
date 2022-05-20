package config

import (
	"Golinker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	err error

	TheDatabase *gorm.DB
)

func Migrations() {
	TheDatabase, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err := TheDatabase.AutoMigrate(&models.LinkModel{})
	if err != nil {
		// migration faiure
		return
	}
}
