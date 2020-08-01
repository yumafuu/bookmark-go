package main

import (
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

	fmt.Println("DELETE LINK")
	fmt.Print("id > ")
	fmt.Scan(&id)
	err := db.Where("id = ?", id).First(&l).Error

	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Delete(&l).Error
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("done")
}
