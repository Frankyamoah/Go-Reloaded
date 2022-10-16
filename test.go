package main

import (
	"fmt"
	"log"
	"os"
	"piscine"
	"strconv"
	"strings"
)

func main() {
	// This reads the contents of the file
	contents, err := os.ReadFile(os.Args[1])
	// Checks for errors
	if err != nil {
		log.Fatal(err)
	}
	// stores contents into string
	stringContent := string(contents)

	// Converts string into string Array
	wordArray := strings.Fields(stringContent)

	// Array of keywords to search through
	// keyWordArray := [8]string{"(hex)", "(bin)", "(up)", "(up,", "(low)", "(low,", "(cap),", "(cap,"}

	// An integer array to store the position of discovered keywords
	var keyWordPosition []int
	// punctuationArray := [6]string{".", ",", "!", "?", ":", ";"}

	// loop through array of arguements
	for i := 0; i < len(wordArray); i++ {
		// swtich statements to cycle through
		switch wordArray[i] {
		case "(hex)":
			strconv.ParseInt(wordArray[i-1], 16, 64)
			keyWordPosition = append(keyWordPosition, i)

		case "(bin)":
			strconv.ParseInt(wordArray[i-1], 2, 64)
			keyWordPosition = append(keyWordPosition, i)

		case "(up)":
			convert := strings.ToUpper(wordArray[i-1])
			keyWordPosition = append(keyWordPosition, i)
			for b := 0; b < len(wordArray); b++ {
				if wordArray[b] == wordArray[i] {
					wordArray = append(wordArray, convert)
				}
			}

		case "(low)":
			convert := strings.ToLower(wordArray[i-1])
			keyWordPosition = append(keyWordPosition, i)
			wordArray = append(wordArray, convert)

		case "(cap)":
			convert := strings.Title(wordArray[i-1])
			keyWordPosition = append(keyWordPosition, i)
			insert(wordArray, i, convert)

		case "(up,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				piscine.ToUpper(wordArray[i-j])
				keyWordPosition = append(keyWordPosition, i)
			}
		case "(low,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				piscine.ToLower(wordArray[i-j])
				keyWordPosition = append(keyWordPosition, i)
			}
		case "(cap,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				piscine.Capitalize(wordArray[i-j])
				keyWordPosition = append(keyWordPosition, i)
			}
		default:
			fmt.Printf(wordArray[i])
		}
	}
	// Deleting the keywords
	counter := 0
	for i := 0; i < len(keyWordPosition); i++ {
		for j := 0; j < len(wordArray); j++ {
			keyWordPosition[i] -= counter
			if i == j {
				wordArray = append(wordArray[:i], wordArray[i:]...)
				counter++

			}
		}
	}
}

func insert(orig []string, index int, value string) []string {
	if index < 0 {
		return nil
	}

	if index >= len(orig) {
		return append(orig, value)
	}
	orig[index] = value
	orig = append(orig, orig[index])

	return orig
}
