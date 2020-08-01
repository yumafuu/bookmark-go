package db

import (
	"bookmarks/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./db.sql"

func NewDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.URL{})
	return db
}
