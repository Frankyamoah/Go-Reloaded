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

	punctuationArray := [6]string{".", ",", "!", "?", ":", ";"}

	// loop through array of arguements
	for i := 0; i < len(wordArray); i++ {
		// loop through array of keyWord
		for j := 0; j < len(keyWordArray); j++ {
			// loop through punctuation Array
			for k := 0; k < len(punctuationArray); k++ {
				for k := 0; k < len(punctuationArray); k++ {
					if (wordArray[i] == keyWordArray[0]) && (keyWordArray[j] == keyWordArray[0]) && (wordArray[i] == keyWordArray[j]) {
					}
				}
			})
		}
	}
}
