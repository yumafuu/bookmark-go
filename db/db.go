package db

import (
	"bookmarks/domain/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const DefaultDbPath = "./db.sql"

func NewDB() *gorm.DB {
	path := os.Getenv("BOOKMARK_GO_DB_PATH")
	if path == "" {
		path = DefaultDbPath
	}

	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.URL{})
	return db
}
