package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset high scores",
	Run:   resetScores,
}

func resetScores(cmd *cobra.Command, args []string) {
	err := os.Remove(getScoresPath())
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error resetting scores:", err)
		return
	}
	fmt.Println("High scores have been reset!")
} 