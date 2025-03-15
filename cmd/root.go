package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wizard",
	Short: "A Terminal User Interface to centralize setup scripts and streamline installations.",
	Long: `Wizard is a setup wizard Terminal User Interface (TUI) application designed to simplify and centralize setup 
or installation scripts. It provides an easy-to-use menu system, powered by Charm Bracelet's Bubble Tea TUI platform,
allowing users to organize and manage scripts all in one place.

With The Setup Wizard, users can:
- Copy and paste shell or programming scripts into its cloud-based database.
- Save these scripts for easy reuse.
- Run any selected script based on their setup or installation requirements.

This tool is perfect for developers and system administrators looking to optimize script execution and management in
a clean, menu-driven terminal interface.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
