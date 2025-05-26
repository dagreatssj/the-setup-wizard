package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a script",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add a script")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
