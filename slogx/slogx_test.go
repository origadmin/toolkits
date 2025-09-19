package slogx

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestDefaultOutputToOutputLog(t *testing.T) {
	// Define the expected default log file name in temp directory
	defaultLogFile := filepath.Join(os.TempDir(), "output.log")

	// Ensure the log file does not exist before the test
	_ = os.Remove(defaultLogFile)

	// Verify that the log file does not exist initially
	_, err := os.Stat(defaultLogFile)
	if !os.IsNotExist(err) {
		t.Fatalf("Expected default log file '%s' to not exist initially, but it does: %v", defaultLogFile, err)
	}

	// Create a new logger without any specific output options
	logger := New(WithOutputFile(defaultLogFile))

	// Log a test message
	testMessage := "This is a test log message for default output."
	logger.Info(testMessage)

	// Verify that the default log file was created
	_, err = os.Stat(defaultLogFile)
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

func TestFileInstanceOutput(t *testing.T) {
	// Define a custom log file name in temp directory
	customLogFile := filepath.Join(os.TempDir(), "custom_output.log")

	// Ensure the log file does not exist before the test
	_ = os.Remove(customLogFile)

	// Verify that the log file does not exist initially
	_, err := os.Stat(customLogFile)
	if !os.IsNotExist(err) {
		t.Fatalf("Expected custom log file '%s' to not exist initially, but it does: %v", customLogFile, err)
	}

	// Create a new logger with a custom output file
	logger := New(WithOutputFile(customLogFile))

	// Log a test message
	testMessage := "This is a test log message for custom output."
	logger.Info(testMessage)

	// Verify that the custom log file was created
	_, err = os.Stat(customLogFile)
	if os.IsNotExist(err) {
		t.Fatalf("Expected custom log file '%s' to be created, but it was not: %v", customLogFile, err)
	} else if err != nil {
		t.Fatalf("Error checking for custom log file '%s': %v", customLogFile, err)
	}

	// Read the content of the log file
	content, err := os.ReadFile(customLogFile)
	if err != nil {
		t.Fatalf("Failed to read custom log file '%s': %v", customLogFile, err)
	}

	// Assert that the log message is present in the file content
	if !strings.Contains(string(content), testMessage) {
		t.Errorf("Expected log file to contain message '%s', but got: %s", testMessage, string(content))
	}

	// Clean up the created log file
	_ = os.Remove(customLogFile)
}

func TestLogLevelOptions(t *testing.T) {
	// Test each log level with isolated log files
	levels := []struct {
		name  string
		level Leveler
	}{
		{"Debug", LevelDebug},
		{"Info", LevelInfo},
		{"Warn", LevelWarn},
		{"Error", LevelError},
	}

	for _, l := range levels {
		t.Run(l.name, func(t *testing.T) {
			// Create unique log file for each test case in temp directory
			logFile := filepath.Join(os.TempDir(), fmt.Sprintf("log_level_%s.log", strings.ToLower(l.name)))
			_ = os.Remove(logFile) // Clean up before test

			logger := New(WithOutputFile(logFile), WithLevel(l.level))
			logger.Debug("Debug message")
			logger.Info("Info message")
			logger.Warn("Warn message")
			logger.Error("Error message")

			content, err := os.ReadFile(logFile)
			if err != nil {
				t.Fatalf("Failed to read log file '%s': %v", logFile, err)
			}

			// Verify only messages at or above the specified level are logged
			switch l.level {
			case LevelDebug:
				if !strings.Contains(string(content), "Debug message") {
					t.Errorf("Expected Debug message in log, but not found")
				}
				if !strings.Contains(string(content), "Info message") {
					t.Errorf("Expected Info message in log, but not found")
				}
				if !strings.Contains(string(content), "Warn message") {
					t.Errorf("Expected Warn message in log, but not found")
				}
				if !strings.Contains(string(content), "Error message") {
					t.Errorf("Expected Error message in log, but not found")
				}
			case LevelInfo:
				if strings.Contains(string(content), "Debug message") {
					t.Errorf("Debug message should not be logged at Info level")
				}
				if !strings.Contains(string(content), "Info message") {
					t.Errorf("Expected Info message in log, but not found")
				}
				if !strings.Contains(string(content), "Warn message") {
					t.Errorf("Expected Warn message in log, but not found")
				}
				if !strings.Contains(string(content), "Error message") {
					t.Errorf("Expected Error message in log, but not found")
				}
			case LevelWarn:
				if strings.Contains(string(content), "Debug message") || strings.Contains(string(content), "Info message") {
					t.Errorf("Debug or Info messages should not be logged at Warn level")
				}
				if !strings.Contains(string(content), "Warn message") {
					t.Errorf("Expected Warn message in log, but not found")
				}
				if !strings.Contains(string(content), "Error message") {
					t.Errorf("Expected Error message in log, but not found")
				}
			case LevelError:
				if strings.Contains(string(content), "Debug message") || strings.Contains(string(content), "Info message") || strings.Contains(string(content), "Warn message") {
					t.Errorf("Debug, Info, or Warn messages should not be logged at Error level")
				}
				if !strings.Contains(string(content), "Error message") {
					t.Errorf("Expected Error message in log, but not found")
				}
			}

			// Clean up after test
			_ = os.Remove(logFile)
		})
	}
}

func TestTimestampOption(t *testing.T) {
	logFile := filepath.Join(os.TempDir(), "timestamp_test.log")
	_ = os.Remove(logFile)

	logger := New(WithOutputFile(logFile), WithTimeLayout(time.DateTime))
	logger.Info("Message with timestamp")

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file '%s': %v", logFile, err)
	}

	// Verify the log contains a valid timestamp in format "2006-01-02 15:04:05"
	t.Logf("Log file content: %s", string(content))
	timestampRegex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}\+\d{2}:\d{2}`)
	if !timestampRegex.MatchString(string(content)) {
		t.Errorf("Expected timestamp in format '2006-01-02 15:04:05', but not found")
	}

	// Clean up
	_ = os.Remove(logFile)
}

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestJSONFormatOption(t *testing.T) {
	logFile := filepath.Join(os.TempDir(), "json_format_test.log")
	_ = os.Remove(logFile)

	logger := New(WithOutputFile(logFile), WithFormat(FormatJSON))
	logger.Info("Message in JSON format", "key", "value")

	content, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Failed to read log file '%s': %v", logFile, err)
	}

	// Verify the log is in JSON format
	if !strings.Contains(string(content), "\"level\":") {
		t.Errorf("Expected JSON format in log, but not found")
	}
	if !strings.Contains(string(content), "\"msg\":") {
		t.Errorf("Expected JSON log to contain message, but not found")
	}
	if !strings.Contains(string(content), "\"key\":\"value\"") {
		t.Errorf("Expected JSON log to contain key-value pair, but not found")
	}

	// Clean up
	_ = os.Remove(logFile)
}
