package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

type difficulty struct {
	name    string
	chances int
}

var difficulties = map[string]difficulty{
	"easy":   {"Easy", 10},
	"medium": {"Medium", 5},
	"hard":   {"Hard", 3},
}

var (
	difficultyFlag string
	maxNumber      int
	// Add random number generator for testing
	randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a new game",
	Run:   runGame,
}

func init() {
	playCmd.Flags().StringVarP(&difficultyFlag, "difficulty", "d", "medium", "Difficulty level (easy|medium|hard)")
	playCmd.Flags().IntVarP(&maxNumber, "max-number", "m", 100, "Maximum number to guess")
}

func runGame(cmd *cobra.Command, args []string) {
	// Add interactive difficulty selection if no flag is provided
	if !cmd.Flags().Changed("difficulty") {
		fmt.Println("\nPlease select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")

		var choice int
		fmt.Print("\nEnter your choice (1-3): ")
		_, err := fmt.Scanf("%d", &choice)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("Invalid choice. Using medium difficulty.")
			difficultyFlag = "medium"
		} else {
			switch choice {
			case 1:
				difficultyFlag = "easy"
			case 2:
				difficultyFlag = "medium"
			case 3:
				difficultyFlag = "hard"
			}
		}
		fmt.Printf("\nSelected %s difficulty!\n", difficultyFlag)
	}

	diff, ok := difficulties[difficultyFlag]
	if !ok {
		fmt.Println("Invalid difficulty level. Using medium.")
		diff = difficulties["medium"]
	}

	fmt.Printf("\nWelcome to the Number Guessing Game!")
	fmt.Printf("\nI'm thinking of a number between 1 and %d", maxNumber)
	fmt.Printf("\nYou have %d chances to guess the correct number.\n\n", diff.chances)

	// Use the injectable random number generator
	target := randGen.Intn(maxNumber) + 1
	chances := diff.chances
	startTime := time.Now()

	for attempts := 1; attempts <= chances; attempts++ {
		var guess int
		fmt.Printf("Enter your guess (%d chances left): ", chances-attempts+1)
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			attempts--
			continue
		}

		if guess == target {
			duration := time.Since(startTime)
			fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts!\n", attempts)
			fmt.Printf("Time taken: %v\n", duration.Round(time.Second))
			updateHighScore(diff.name, attempts, duration)
			return
		}

		if guess < target {
			fmt.Println("The number is greater than your guess.")
		} else {
			fmt.Println("The number is less than your guess.")
		}
	}

	fmt.Printf("\nGame Over! The number was %d\n", target)
}
