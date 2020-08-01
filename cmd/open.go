package cmd

import (
	"bookmarks/cli"
	"bookmarks/db"
	"bookmarks/domain/models"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open url",
	Long:  `open url`,
	Run:   open,
}

func open(cmd *cobra.Command, args []string) {
	db := db.NewDB()
	defer db.Close()

	var u models.URL
	var s string

	cli.GetInput("id", &s)
	items := strings.Fields(s)

	if len(items) == 0 {
		return
	}

	err := db.Where("id = ?", items[0]).First(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("open %v\n", u.Url)
	exec.Command("open", u.Url).Run()

	u.LastVisitAt = time.Now()
	err = db.Save(&u).Error
	if err != nil {
		log.Fatal(err)
	}
}
