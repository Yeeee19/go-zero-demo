package database

import (
	"gin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func Setup() {
	db, err := gorm.Open(sqlite.Open("jpa.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})

	gdb = db
}

func GetDB() *gorm.DB {
	if gdb == nil {
		panic("database no setup")
	}

	return gdb
}
