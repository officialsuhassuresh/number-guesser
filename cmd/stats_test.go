package cmd

import (
	"strings"
	"testing"
	"time"
)

func TestHighScores(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	// Test updating and loading high scores
	t.Run("Update and load high scores", func(t *testing.T) {
		score := Score{
			Attempts: 3,
			Duration: 5 * time.Second,
			Date:     time.Now(),
		}

		// Update score for medium difficulty
		updateHighScore("Medium", score.Attempts, score.Duration)

		// Load scores and verify
		scores := loadHighScores()
		if scores.Medium.Attempts != score.Attempts {
			t.Errorf("expected attempts %d, got %d", score.Attempts, scores.Medium.Attempts)
		}
	})

	// Test score display
	t.Run("Display scores", func(t *testing.T) {
		output := captureOutput(func() {
			showStats(nil, nil)
		})

		expectedOutputs := []string{
			"High Scores",
			"Easy: No high score yet",
			"Medium",
			"Hard: No high score yet",
		}

		for _, expected := range expectedOutputs {
			if !strings.Contains(output, expected) {
				t.Errorf("expected output to contain %q, got %q", expected, output)
			}
		}
	})
}
