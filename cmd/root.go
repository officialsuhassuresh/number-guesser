package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "number-guesser",
	Short:   "A number guessing game",
	Version: version,
	Long: `A fun CLI number guessing game where you try to guess 
a randomly generated number within a limited number of attempts.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(resetCmd)
} 