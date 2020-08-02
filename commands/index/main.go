package main

import (
	"bookmarks/db"
	"bookmarks/domain/models"
	"fmt"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	var urls []models.URL
	db.Order("last_visit_at desc").Find(&urls)

	for _, u := range urls {
		fmt.Println(u.ID, u.KeyWord, u.Url, u.Title)
	}
}
