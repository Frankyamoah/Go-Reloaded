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

	// Variable of converted input
	convText := converter(wordArray)

	// Seperates text by spaces
	finalText := strings.Join(convText, " ")

	// Creation of the result file where the converted text is displayed
	result, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	// closing of newly created file
	defer result.Close()

	// loops through final string writing it into the file
	for _, word := range finalText {
		_, err := result.WriteString(string(word))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// loop through array of arguements
func converter(wordArray []string) []string {
	var result []string
	for i := 0; i < len(wordArray); i++ {
		// swtich statements to cycle through

		switch wordArray[i] {
		case "(hex)":
			convHex, _ := strconv.ParseInt(wordArray[i-1], 16, 64)
			result[len(result)-1] = fmt.Sprint(convHex)

		case "(bin)":
			convBin, _ := strconv.ParseInt(wordArray[i-1], 2, 64)
			result[len(result)-1] = fmt.Sprint(convBin)

		case "(up)":
			convUP := strings.ToUpper(wordArray[i-1])
			result[len(result)-1] = fmt.Sprint(convUP)

		case "(low)":
			convLOW := strings.ToLower(wordArray[i-1])
			result[len(result)-1] = fmt.Sprint(convLOW)

		case "(cap)":
			convCAP := strings.Title(wordArray[i-1])
			result[len(result)-1] = fmt.Sprint(convCAP)

		case "(up,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				convUP2 := strings.ToUpper(result[len(result)-j])
				result[len(result)-j] = fmt.Sprint(convUP2)
			}
			i++
		case "(low,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				convLOW2 := strings.ToLower(result[len(result)-j])
				result[len(result)-j] = fmt.Sprint(convLOW2)
			}
			i++
		case "(cap,":
			number := strings.Trim(wordArray[i+1], ")")
			integer, _ := strconv.Atoi(number)
			for j := 1; j <= integer; j++ {
				convCAP2 := strings.Title(result[len(result)-j])
				result[len(result)-j] = fmt.Sprint(convCAP2)

			}
			i++
		default:
			result = append(result, wordArray[i])
		}
	}
	return puncEditor(vowel(result))
}

func puncEditor(wordArr []string) []string {
	punctuationArray := [6]string{".", ",", "!", "?", ":", ";"}
	for i, word := range wordArr {
		for _, punc := range punctuationArray {
			// for dealing with punc connecting to word
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				wordArr[i-1] += punc
				wordArr[i] = word[1:]
			}
		}
	}
	// for dealing with punc at end of string
	for i, word := range wordArr {
		for _, punc := range punctuationArray {
			if (string(word[0]) == punc) && (wordArr[len(wordArr)-1] == wordArr[i]) {
				wordArr[i-1] += word
				wordArr = wordArr[:len(wordArr)-1]
			}
		}
	}
	// for dealing with punc in the middle of string
	for i, word := range wordArr {
		for _, punc := range punctuationArray {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && wordArr[i] != wordArr[len(wordArr)-1] {
				wordArr[i-1] += word
				wordArr = append(wordArr[:i], wordArr[i+1:]...)
			}
		}
	}
	// If apostrophe is before a word this joins it to the next word
	count := 0
	for i, word := range wordArr {
		if word == "'" && count == 0 {
			count += 1
			wordArr[i+1] = word + wordArr[i+1]
			wordArr = append(wordArr[:i], wordArr[i+1:]...)
		}
	}
	// This connects second apostraphe to the end of word
	for i, word := range wordArr {
		if word == "'" {
			wordArr[i-1] = wordArr[i-1] + word
			wordArr = append(wordArr[:i], wordArr[i+1:]...)
		}
	}

	return wordArr
}

func vowel(wordArr []string) []string {
	// Array of Vowels
	vowels := [12]string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	// Loop through string Array
	for i := range wordArr {
		// Loop through punctuation Array
		for _, vowel := range vowels {
			if wordArr[i] == "a" {
				// The next word is placed in variable
				variable := []byte(wordArr[i+1])
				for x := 0; x < len(variable); x++ {
					// refers to the first letter in string to see if it is a voewl
					if string(variable[0]) == vowel {
						// If true add n to the a
						wordArr[i] += "n"
						break

					}
				}
			}
		}
	}
	return wordArr
}
