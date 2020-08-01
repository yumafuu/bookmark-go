package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./db.sql"

var Conn *sql.DB

type Link struct {
	id          int
	uid         string
	title       string
	url         string
	visit_count int
}

func init() {
	Conn, err := sql.Open("sqlite3", dbPath)
	defer Conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(dbPath)
	if err != nil {
		os.Create(dbPath)
	}

	cmd := `CREATE TABLE IF NOT EXISTS links(
             id INT NOT NULL PRIMARY KEY,
						 uid STRING NOT NULL,
             title STRING NOT NULL,
						 url STRING NOT NULL,
						 visit_count INT NOT NULL DEFAULT 0
					 )`

	_, err = Conn.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	Conn, _ := sql.Open("sqlite3", dbPath)
	defer Conn.Close()

	q := "SELECT * FROM links order by visit_count desc"
	rows, _ := Conn.Query(q)
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var l Link
		err := rows.Scan(&l.id, &l.uid, &l.title, &l.url, &l.visit_count)
		if err != nil {
			log.Println(err)
		}

		links = append(links, l)
	}

	for _, l := range links {
		fmt.Println(l.uid, l.url, l.title)
	}
}
