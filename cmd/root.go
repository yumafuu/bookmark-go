package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bookmark-go",
	Short: "bookmark-go is a bookmark tool on cli",
	Long:  `bookmark-go is a bookmark tool on cli`,
	Run:   indexCmd.Run,
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
