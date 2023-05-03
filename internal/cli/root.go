package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	node   string
	signer string
)

var rootCmd = &cobra.Command{
	Use:   "bounty",
	Short: "the bounty command-line interface",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&node, "node", "https://eth.llamarpc.com", "set the `URL` of the execution layer node")
	rootCmd.PersistentFlags().StringVar(&signer, "signer", "http://localhost:8550", "set the `URL` of the clef external signer")

	rootCmd.AddCommand(openCmd)
}

func Run() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
