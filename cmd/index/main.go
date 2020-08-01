package main

import (
	"bookmarks/db"
	"bookmarks/models"
	"fmt"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	var urls []models.URL
	db.Find(&urls)

	fmt.Println("ID KeyWord URL TITLE")
	for _, u := range urls {
		fmt.Println(u.ID, u.KeyWord, u.Url, u.Title)
	}
}
