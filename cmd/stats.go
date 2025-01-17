package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

type Score struct {
	Attempts int           `json:"attempts"`
	Duration time.Duration `json:"duration"`
	Date     time.Time     `json:"date"`
}

type HighScores struct {
	Easy   Score `json:"easy"`
	Medium Score `json:"medium"`
	Hard   Score `json:"hard"`
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show high scores",
	Run:   showStats,
}

var getScoresPath = func() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "scores.json"
	}
	return filepath.Join(homeDir, ".number-guesser-scores.json")
}

func showStats(cmd *cobra.Command, args []string) {
	scores := loadHighScores()
	fmt.Println("\nHigh Scores:")
	fmt.Println("------------")
	printScore("Easy", scores.Easy)
	printScore("Medium", scores.Medium)
	printScore("Hard", scores.Hard)
}

func printScore(difficulty string, score Score) {
	if score.Attempts > 0 {
		fmt.Printf("%s: %d attempts in %v (achieved on %s)\n",
			difficulty, score.Attempts, score.Duration.Round(time.Second),
			score.Date.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("%s: No high score yet\n", difficulty)
	}
}

func loadHighScores() HighScores {
	var scores HighScores
	data, err := os.ReadFile(getScoresPath())
	if err != nil {
		return scores
	}
	json.Unmarshal(data, &scores)
	return scores
}

func updateHighScore(difficulty string, attempts int, duration time.Duration) {
	scores := loadHighScores()
	newScore := Score{
		Attempts: attempts,
		Duration: duration,
		Date:     time.Now(),
	}

	switch difficulty {
	case "Easy":
		if scores.Easy.Attempts == 0 || attempts < scores.Easy.Attempts {
			scores.Easy = newScore
		}
	case "Medium":
		if scores.Medium.Attempts == 0 || attempts < scores.Medium.Attempts {
			scores.Medium = newScore
		}
	case "Hard":
		if scores.Hard.Attempts == 0 || attempts < scores.Hard.Attempts {
			scores.Hard = newScore
		}
	}

	data, _ := json.Marshal(scores)
	os.WriteFile(getScoresPath(), data, 0644)
}
