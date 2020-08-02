package cmd

import (
	"bookmarks/cli"
	"bookmarks/db"
	"bookmarks/domain/models"
	"fmt"
	"log"

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
