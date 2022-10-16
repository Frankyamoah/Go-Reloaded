package main

import (
	"fmt"
	"log"
	"os"
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
	converter(wordArray)
}

// func insert(orig []string, index int, value string) []string {
// 	if index < 0 {
// 		return nil
// 	}

// 	if index >= len(orig) {
// 		return append(orig, value)
// 	}
// 	orig[index] = value
// 	orig = append(orig, orig[index])

// 	return orig
// }

// Array of keywords to search through
// keyWordArray := [8]string{"(hex)", "(bin)", "(up)", "(up,", "(low)", "(low,", "(cap),", "(cap,"}

// An integer array to store the position of discovered keywords
var keyWordPosition []int

// punctuationArray := [6]string{".", ",", "!", "?", ":", ";"}

// loop through array of arguements
func converter(wordArray []string) []string {
	var result []string
	for i := 0; i < len(wordArray); i++ {
		// swtich statements to cycle through

		switch wordArray[i] {
		case "(hex)":
			convHex, _ := strconv.ParseInt(wordArray[i-1], 16, 64)
			keyWordPosition = append(keyWordPosition, i)
			result[len(result)-1] = fmt.Sprint(convHex)

		case "(bin)":
			convBin, _ := strconv.ParseInt(wordArray[i-1], 2, 64)
			keyWordPosition = append(keyWordPosition, i)
			result[len(result)-1] = fmt.Sprint(convBin)

		case "(up)":
			convUP := strings.ToUpper(wordArray[i-1])
			keyWordPosition = append(keyWordPosition, i)
			result[len(result)-1] = fmt.Sprint(convUP)

		case "(low)":
			convLOW := strings.ToLower(wordArray[i-1])
			keyWordPosition = append(keyWordPosition, i)
			result[len(result)-1] = fmt.Sprint(convLOW)

		case "(cap)":
			convCAP := strings.Title(wordArray[i-1])
			keyWordPosition = append(keyWordPosition, i)
			result[len(result)-1] = fmt.Sprint(convCAP)

		case "(up,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				convUP2 := strings.ToUpper(result[len(result)-j])
				result[len(result)-j] = fmt.Sprint(convUP2)
				keyWordPosition = append(keyWordPosition, i)
			}
		case "(low,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				convLOW2 := strings.ToLower(result[len(result)-j])
				result[len(result)-j] = fmt.Sprint(convLOW2)
				keyWordPosition = append(keyWordPosition, i)
			}
		case "(cap,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				convCAP2 := strings.Title(result[len(result)-j])
				result[len(result)-j] = fmt.Sprint(convCAP2)
				keyWordPosition = append(keyWordPosition, i)
			}
		default:
			result = append(result, wordArray[i])
		}
	}
	// Deleting the keywords
	counter := 0
	for i := 0; i < len(keyWordPosition); i++ {
		for j := 0; j < len(result); j++ {
			keyWordPosition[i] -= counter
			if i == j {
				result = append(result[:i], result[i:]...)
				counter++

			}
		}
	}
	return result
}
