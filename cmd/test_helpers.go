package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// testSetup creates a temporary scores file for testing
func testSetup(t *testing.T) (string, func()) {
	t.Helper()
	tmpDir, err := os.MkdirTemp("", "number-guesser-test")
	if err != nil {
		t.Fatal(err)
	}

	// Store the original function
	originalGetScoresPath := getScoresPath

	// Replace with test version
	getScoresPath = func() string {
		return filepath.Join(tmpDir, "test-scores.json")
	}

	// Return cleanup function
	return tmpDir, func() {
		getScoresPath = originalGetScoresPath
		os.RemoveAll(tmpDir)
	}
}

// captureOutput captures stdout during test execution
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
