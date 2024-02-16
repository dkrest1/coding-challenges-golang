package main

import (
	"bufio"
	"fmt"
	"os"
)


func byteCount(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	fileSize := fileInfo.Size()

	return fileSize, nil
}

func countLine(filename string) (int, error) {
	file, err := os.Open(filename)
	
	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}