package cmd

import (
	"bookmarks/db"
	"bookmarks/domain/models"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete url",
	Long:  `delete url`,
	Run:   delete,
}

func delete(cmd *cobra.Command, args []string) {
	db := db.NewDB()
	defer db.Close()

	var urls []models.URL
	db.Order("last_visit_at desc").Find(&urls)

	for _, u := range urls {
		fmt.Println(u.ID, u.KeyWord, u.Url, u.Title)
	}

}
