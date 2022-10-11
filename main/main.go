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

	// An integer array to store the position of discovered keywords
	var keyWordPosition []int
	// punctuationArray := [6]string{".", ",", "!", "?", ":", ";"}

	// loop through array of arguements
	for i := 0; i < len(wordArray); i++ {
		// loop through array of keyWord
		for j := 0; j < len(keyWordArray); j++ {
			// swtich statements to cycle through
			switch wordArray[i] {
			case "(hex)":
				piscine.convertHexToDecimal(wordArray[i-1])
				keyWordPosition = append(keyWordPosition, i)

			case "(bin)":
				piscine.convertBinaryToDecimal(wordArray[i-1])
				keyWordPosition = append(keyWordPosition, i)

			case "(up)":
				piscine.ToUpper(wordArray[i-1])
				keyWordPosition = append(keyWordPosition, i)

			case "(low)":
				piscine.ToLower(wordArray[i-1])
				keyWordPosition = append(keyWordPosition, i)

			case "(cap)":
				piscine.Capitalize(wordArray[i-1])
				keyWordPosition = append(keyWordPosition, i)

			}
		}
	}
}
