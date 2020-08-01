package main

import (
	"bookmarks/db"
	"bookmarks/models"
	"fmt"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	var keyWord string
	var url string

	fmt.Println("ADD LINK")
	fmt.Print("keyWord > ")
	fmt.Scan(&keyWord)
	fmt.Print("url > ")
	fmt.Scan(&url)

	l, err := models.NewURL(
		keyWord, url,
	)

	if err != nil {
		fmt.Printf("FAIL: %v\n", err)
		return
	}

	err = db.Create(&l).Error
	if err != nil {
		fmt.Printf("FAIL: %v\n", err)
		return
	}

	fmt.Println("Add link successfully")
}
