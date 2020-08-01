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
	var l models.URL
	var id string

	fmt.Println("DELETE URL")
	cli.GetInput("id", &id)

	err := db.Where("id = ?", id).First(&l).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Delete(&l).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
