package main

import (
	"bookmarks/cli"
	"bookmarks/db"
	"bookmarks/domain/models"
	"fmt"
	"log"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	var keyWord string
	var url string

	fmt.Println("ADD LINK")
	cli.WaitingInput("keyWord", &keyWord)
	cli.WaitingInput("URL", &url)

	l, err := models.NewURL(
		keyWord, url,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&l).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Add link successfully")
}
