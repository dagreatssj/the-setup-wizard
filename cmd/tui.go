package cmd

import (
	"fmt"
	"github.com/dagreatssj/the-setup-wizard/internal/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the TUI application",
	Run: func(cmd *cobra.Command, args []string) {
		err := tui.RunTeaApp()
		if err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
