package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./db.sql"

var DbConnection *sql.DB

type Link struct {
	id    int
	title string
	url   string
	count int
}

func init() {
	_, err := os.Stat(dbPath)
	if err != nil {
		os.Create(dbPath)
	}
}

func main() {
	fmt.Println(dbPath)
	conn, err := sql.Open("sqlite3", dbPath)
	defer conn.Close()

	if err != nil {
		log.Println("---------------------------")
		log.Fatal(err)
	}

	q := "SELECT * FROM links order by count desc"

	_, err = conn.Exec(q)
	if err != nil {
		log.Fatal(err)
	}

}
