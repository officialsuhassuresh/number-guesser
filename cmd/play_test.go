package cmd

import (
	"testing"
)

func TestDifficulties(t *testing.T) {
	expected := map[string]int{
		"easy":   10,
		"medium": 5,
		"hard":   3,
	}

	for name, diff := range difficulties {
		if diff.chances != expected[name] {
			t.Errorf("difficulty %q: expected %d chances, got %d", name, expected[name], diff.chances)
		}
	}
}
