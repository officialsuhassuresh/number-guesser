package cmd

import (
	"os"
	"testing"
	"time"
)

func TestResetScores(t *testing.T) {
	_, cleanup := testSetup(t)
	defer cleanup()

	// Create a test score file
	score := Score{
		Attempts: 3,
		Duration: 5 * time.Second,
		Date:     time.Now(),
	}
	updateHighScore("Medium", score.Attempts, score.Duration)

	// Test reset functionality
	t.Run("Reset scores", func(t *testing.T) {
		resetScores(nil, nil)

		// Verify scores file is removed
		if _, err := os.Stat(getScoresPath()); !os.IsNotExist(err) {
			t.Error("expected scores file to be removed")
		}

		// Verify loading scores returns empty scores
		scores := loadHighScores()
		if scores.Medium.Attempts != 0 {
			t.Errorf("expected empty scores, got attempts: %d", scores.Medium.Attempts)
		}
	})
}
