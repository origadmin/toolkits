package slogx

import (
	"os"
	"strings"
	"testing"
)

func TestDefaultOutputToOutputLog(t *testing.T) {
	// Define the expected default log file name
	defaultLogFile := "output.log"

	// Ensure the log file does not exist before the test
	_ = os.Remove(defaultLogFile)

	// Create a new logger without any specific output options
	logger := New()

	// Log a test message
	testMessage := "This is a test log message for default output."
	logger.Info(testMessage)

	// Verify that the default log file was created
	_, err := os.Stat(defaultLogFile)
	if os.IsNotExist(err) {
		t.Fatalf("Expected default log file '%s' to be created, but it was not: %v", defaultLogFile, err)
	} else if err != nil {
		t.Fatalf("Error checking for default log file '%s': %v", defaultLogFile, err)
	}

	// Read the content of the log file
	content, err := os.ReadFile(defaultLogFile)
	if err != nil {
		t.Fatalf("Failed to read default log file '%s': %v", defaultLogFile, err)
	}

	// Assert that the log message is present in the file content
	if !strings.Contains(string(content), testMessage) {
		t.Errorf("Expected log file to contain message '%s', but got: %s", testMessage, string(content))
	}

	// Clean up the created log file
	_ = os.Remove(defaultLogFile)
}
