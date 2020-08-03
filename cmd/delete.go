package cmd

import (
	"bookmark-go/cli"
	"bookmark-go/db"
	"bookmark-go/domain/models"
	"fmt"
	"log"
	"strings"

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

	var u models.URL
	var s string

	fmt.Println("DELETE URL")
	if id == "" {
		cli.GetInput("id", &s)
		items := strings.Fields(s)
		if len(items) == 0 {
			return
		}
		id = items[0]
	}

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Delete(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("delete %v %v\n", u.Title, u.Url)
}
