package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define command-line options
	countBytes := flag.Bool("c", false, "Count bytes")
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countChars := flag.Bool("m", false, "Count characters")
	flag.Parse()

	filename := flag.Arg(0)
	fmt.Println(filename)

	// If no filename is provided, read from standard input
	if filename == "" {
		// Implement reading from standard input if needed
		fmt.Println("Reading from standard input is not yet implemented.")
		return
	}

	// Check if at least one count option is provided
	if !(*countBytes || *countLines || *countWords || *countChars) {
		*countBytes = true
		*countLines = true
		*countWords = true
	}

	// Call the appropriate functions based on options
	if *countBytes {
		bytesCount, err := byteCount(filename)
		if err != nil {
			fmt.Println("Error counting bytes:", err)
			return
		}
		fmt.Printf("%8d %s\n", bytesCount, filename)
	}

	if *countLines {
		linesCount, err := lineCount(filename)
		if err != nil {
			fmt.Println("Error counting lines:", err)
			return
		}
		fmt.Printf("%8d %s\n", linesCount, filename)
	}

	if *countWords {
		wordsCount, err := wordCount(filename)
		if err != nil {
			fmt.Println("Error counting words:", err)
			return
		}
		fmt.Printf("%8d %s\n", wordsCount, filename)
	}

	if *countChars {
		charsCount, err := charCount(filename)
		if err != nil {
			fmt.Println("Error counting characters:", err)
			return
		}
		fmt.Printf("%8d %s\n", charsCount, filename)
	}
}

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
		fmt.Println("Error", err)
		return 0, err
	}

	return wordCount, nil
}

func charCount(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	charCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}

	return charCount, nil
} 