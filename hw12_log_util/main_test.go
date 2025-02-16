package main

import (
	"os"
	"testing"
)

func TestAnalyzeStatistic(t *testing.T) {
	logs := `
info | Test log 1
error | Test log 2
info | Test log 3
warning | Test log 4
info | Test log 5
`
	tmpFile, err := os.CreateTemp("", "testing*.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, rErr := tmpFile.WriteString(logs); rErr != nil {
		t.Fatalf("Failed to write to temp file: %v", rErr)
	}
	tmpFile.Close()

	levelLog := "info"
	result, err := analyzeStatistic(tmpFile.Name(), levelLog)
	if err != nil {
		t.Fatalf("analyzeStatistic returned an error: %v", err)
	}

	expectedCount := 3
	if result[levelLog] != expectedCount {
		t.Errorf("Expected %d info logs, got %d", expectedCount, result[levelLog])
	}
}

func TestWriteOutputToFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "output.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	result := map[string]int{"info": 3}

	err = writeOutput(tmpFile.Name(), result)
	if err != nil {
		t.Fatalf("writeOutput returned an error: %v", err)
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	expectedContent := "Log level: info | Count: 3\n"
	if string(content) != expectedContent {
		t.Errorf("Expected file content to be %q, got %q", expectedContent, string(content))
	}
}

func TestWriteOutputToConsole(t *testing.T) {
	result := map[string]int{"info": 3}

	err := writeOutput("", result)
	if err != nil {
		t.Fatalf("writeOutput returned an error: %v", err)
	}
}
