package main

import (
	"fmt"
	"os"
	"testing"
)

func TestByteCount(t *testing.T) {
	// Create a temporary file for testing
	tempFile := createTempFile("test content")
	defer cleanupTempFile(tempFile)

	// Call the function being tested
	result, err := byteCount(tempFile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedSize := int64(12)
	expectedError := error(nil) 

	// Check if the result matches expectations
	if result != expectedSize {
		t.Errorf("Expected size %d, but got %d", expectedSize, result)
	}

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}

func TestLineCount(t *testing.T) {
	// Create a temporary file for testing
	tempFile := createTempFile("line1\nline2\nline3\n")
	defer cleanupTempFile(tempFile)

	// Call the function being tested
	result, err := lineCount(tempFile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedCount := 3
	expectedError := error(nil)

	// Check if the result matches expectations
	if result != expectedCount {
		t.Errorf("Expected line count %d, but got %d", expectedCount, result)
	}

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}

func TestWordCount(t *testing.T) {
	// Create a temporary file for testing
	tempFile := createTempFile("word1 word2 word3")
	defer cleanupTempFile(tempFile)

	// Call the function being tested
	result, err := wordCount(tempFile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedCount := 3
	expectedError := error(nil)

	// Check if the result matches expectations
	if result != expectedCount {
		t.Errorf("Expected word count %d, but got %d", expectedCount, result)
	}

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}

func TestCharCount(t *testing.T) {
	// Create a temporary file for testing
	tempFile := createTempFile("abcd1234!@#$")
	defer cleanupTempFile(tempFile)

	// Call the function being tested
	result, err := charCount(tempFile)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedCount := 12
	expectedError := error(nil)

	// Check if the result matches expectations
	if result != expectedCount {
		t.Errorf("Expected character count %d, but got %d", expectedCount, result)
	}

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}

// Helper function to create a temporary file for testing
func createTempFile(content string) *os.File {
	file, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
	}

	file.WriteString(content)
	return file
}

// Helper function to clean up the temporary file after testing
func cleanupTempFile(file *os.File) {
	if err := os.Remove(file.Name()); err != nil {
		fmt.Println("Error removing temp file:", err)
	}
}
