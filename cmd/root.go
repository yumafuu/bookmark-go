package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	keyWord string
	url     string
	id      string
	rootCmd = &cobra.Command{
		Use:   "bookmark-go",
		Short: "bookmark-go is a bookmark tool on cli",
		Long:  `bookmark-go is a bookmark tool on cli`,
		Run:   indexCmd.Run,
	}
)

func init() {
	addCmd.PersistentFlags().StringVarP(&keyWord, "keyWord", "k", "", "add keyWord")
	addCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "open this link")
	openCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "specify id")
	deleteCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "specify id")
}

func Execute() {
	all := []*cobra.Command{
		addCmd,
		indexCmd,
		deleteCmd,
		openCmd,
	}

	for _, c := range all {
		rootCmd.AddCommand(c)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
