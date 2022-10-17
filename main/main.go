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

	fmt.Println(converter(wordArray))
}

func vowel(wordArr []string) []string {
	vowels := [12]string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i, _ := range wordArr {
		for _, vowel := range vowels {
			if wordArr[i] == "a" {
				variable := []byte(wordArr[i+1])
				for x := 0; x < len(variable); x++ {
					if string(variable[0]) == vowel {
						wordArr[i] += "n"
						break

					}
				}
			}
		}
	}
	return wordArr
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
	count := 0
	for i, word := range wordArr {
		if word == "'" && count == 0 {
			count += 1
			wordArr[i+1] = word + wordArr[i+1]
			wordArr = append(wordArr[:i], wordArr[i+1:]...)
		}
	}
	for i, word := range wordArr {
		if word == "'" {
			wordArr[i-1] = wordArr[i-1] + word
			wordArr = append(wordArr[:i], wordArr[i+1:]...)
		}
	}

	return wordArr
}

// bob
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
