package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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

	var filename *os.File

	if flag.NArg() == 0 {
		stat, err := os.Stdin.Stat()
		if err != nil {
			log.Println("Error:", err)
			return
		}

		if (stat.Mode() & os.ModeCharDevice) == 0 {
			filename = os.Stdin
		} else {
			log.Println("Reading from standard input is not yet implemented.")
			return
		}
	} else {
		filenamePath := flag.Arg(0)
		file, err := os.Open(filenamePath)
		if err != nil {
			log.Println("Error:", err)
			return
		}
		defer file.Close()

		filename = file
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
			log.Println("Error counting bytes:", err)
			return
		}
		fmt.Printf("%8d %v\n", bytesCount, filename.Name())
	}

	if *countLines {
		linesCount, err := lineCount(filename)
		if err != nil {
			log.Println("Error counting lines:", err)
			return
		}
		fmt.Printf("%8d %v\n", linesCount, filename.Name())
	}

	if *countWords {
		wordsCount, err := wordCount(filename)
		if err != nil {
			log.Println("Error counting words:", err)
			return
		}
		fmt.Printf("%8d %v\n", wordsCount, filename.Name())
	}

	if *countChars {
		charsCount, err := charCount(filename)
		if err != nil {
			log.Println("Error counting characters:", err)
			return
		}
		fmt.Printf("%8d %v\n", charsCount, filename.Name())
	}
}

func byteCount(filename *os.File) (int64, error) {
	fileInfo, err := filename.Stat()
	if err != nil {
		log.Println("Error:", err)
		return 0, err
	}

	fileSize := fileInfo.Size()

	return fileSize, nil
}

func lineCount(filename *os.File) (int, error) {
	filename.Seek(0, 0) 
	scanner := bufio.NewScanner(filename)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error:", err)
		return 0, err
	}

	return lineCount, nil
}

func wordCount(filename *os.File) (int, error) {
	filename.Seek(0, 0) 
	scanner := bufio.NewScanner(filename)
	wordCount := 0


	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error", err)
		return 0, err
	}

	return wordCount, nil
}

func charCount(filename *os.File) (int, error) {
	filename.Seek(0, 0) 
	scanner := bufio.NewScanner(filename)
	charCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		charCount += len(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error: ", err)
		return 0, err
	}

	return charCount, nil
} 