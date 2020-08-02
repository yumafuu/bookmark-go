package cmd

import (
	"bookmarks/db"
	"bookmarks/domain/models"
	"fmt"

	"github.com/spf13/cobra"
)

var indexCmd = &cobra.Command{
	Use:   "show list",
	Short: "show url list",
	Long:  `show url list`,
	Run:   index,
}

func index(cmd *cobra.Command, args []string) {
	db := db.NewDB()
	defer db.Close()

	var urls []models.URL
	db.Order("last_visit_at desc").Find(&urls)

	for _, u := range urls {
		fmt.Println(u.ID, u.KeyWord, u.Url, u.Title)
	}
}
