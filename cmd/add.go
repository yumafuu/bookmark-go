package cmd

import (
	"bookmark-go/cli"
	"bookmark-go/db"
	"bookmark-go/domain/models"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add url",
	Long:  `add url`,
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	db := db.NewDB()
	defer db.Close()

	fmt.Println("ADD URL")
	if keyWord == "" {
		cli.GetInput("keyWord", &keyWord)
	}

	if url == "" {
		cli.GetInput("URL", &url)
	}

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

	fmt.Println("Add url successfully")
}
