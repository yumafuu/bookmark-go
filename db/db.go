package db

import (
	"bookmarks/domain/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const DefaultDbFile = "/.bookmark.db"

func NewDB() *gorm.DB {
	path, err := dbPath()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(path)
	if err != nil {
		_, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(path)

	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.URL{})
	return db
}

func dbPath() (string, error) {
	path := os.Getenv("BOOKMARK_GO_DB_PATH")
	if path != "" {
		return path, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}
	path = home + DefaultDbFile

	return path, nil
}
