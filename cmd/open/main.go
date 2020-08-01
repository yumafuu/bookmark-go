package main

import (
	"bookmarks/db"
	"bookmarks/domain/models"
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

	fmt.Print("id > ")
	i := getInput()
	inputs := strings.Fields(i)

	if len(inputs) == 0 {
		return
	}
	id := inputs[0]
	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Url())
	exec.Command("open", u.Url()).Run()
}

func getInput() string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}
