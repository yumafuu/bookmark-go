package main

import (
	"bookmarks/db"
	"bookmarks/models"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	var u models.URL

	i := getInput()
	items := strings.Fields(i)

	err := db.Where("id = ?", items[0]).First(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(items[1])
	exec.Command("open", items[1]).Run()
}

func getInput() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}
