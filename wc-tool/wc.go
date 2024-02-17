package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func lineCount(filename string) (int, error) {
	file, err := os.Open(filename)
	
	if err != nil {
		fmt.Println("Error:", err)
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

func wordCount(filename string) (int, error) {
	file, err :=  os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		// fmt.Println('Error', err)
		return 0, err
	}

	return wordCount, nil
}