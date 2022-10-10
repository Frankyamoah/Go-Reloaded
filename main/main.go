package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// This reads the contents of the file
	contents, err := os.ReadFile(os.Args[1])
	// Checks for errors
	if err != nil {
		log.Fatal(err)
	}
	// Converts arguements into string
	stringContent := string(contents)

	// Converts string into string Array
	wordArray := strings.Fields(stringContent)

	// Array of keywords to search through
	keyWordArray := [8]string{"(hex)", "(bin)", "(up)", "(up,", "(low)", "(low,", "(cap),", "(cap,"}

	// loop through array of arguements

	for i := 0; i < len(wordArray); i++ {
		// loop through array of keyWords
		for j := 0; j < len(keyWordArray); j++ {
		}
	}
}
