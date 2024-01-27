package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("dummyFile.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error in openning the file!", err)
		return
	}

	lineCount := 0
	wordCount := 0
	charCount := 0
	whitespace := 0
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			// fmt.Println(size, err)

			//EOF is returned as an error
			if err == io.EOF {
				lineCount++
			}

			break
		}

		if r == '\n' {
			lineCount += 1
			charCount++
		} else if r == ' ' {
			whitespace++
			wordCount++
			charCount++
		} else {
			charCount++
			// fmt.Printf("Read rune: %c, Size: %d\n", r, size)
		}
	}

	fileStats, _ := file.Stat()
	byteLength := fileStats.Size()
	fmt.Printf("Lines in file: %d\nWords in file %d\nCharacters in file: %d\nBytes in file: %d \n", lineCount, lineCount+wordCount, charCount, byteLength)
}
