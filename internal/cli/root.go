package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bounty",
	Short: "the bounty command-line interface",
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
