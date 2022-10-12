package piscine

import (
	"strconv"
)

// Capitilize function
func Capitalize(s string) string {
	empty := ""
	stringRune := []rune(s)
	for i := 0; i < len(stringRune); i++ {
		a := string(stringRune[i])
		// if 'i' is not at index 0 and the letter before i is alphanumerical, change it to lowercase
		if i != 0 && IsAlpha(string(stringRune[i-1])) {
			empty += ToLower(a)
			// if letter before 'i' not a alphanumerical, change it to uppercase
		} else {
			empty += ToUpper(a)
		}
	}
	return empty
}

// function to check if element is capital letter, neccesary to utilize capitilize
func IsAlpha(s string) bool {
	stringRune := []rune(s)
	for i := 0; i < len(s); i++ {
		a := stringRune[i]
		if a >= 32 && a < 48 {
			return false
		}
		if a >= 58 && a < 65 {
			return false
		}
		if a >= 91 && a < 97 {
			return false
		}
		if a >= 123 {
			return false
		}

	}
	return true
}

// function fo checking if element is lowercase

func IsLower(s string) bool {
	stringRune := []rune(s)
	for i := 0; i < len(s); i++ {
		if stringRune[i] <= 96 {
			return false
		}
		if stringRune[i] >= 123 {
			return false
		}
	}
	return true
}

// checks if element is numerical
func IsNumeric(s string) bool {
	stringRune := []rune(s)
	for i := 0; i < len(s); i++ {
		if !(stringRune[i] >= 48 && stringRune[i] <= 57) {
			return false
		}
	}
	return true
}

// changes element to uppercase

func ToUpper(s string) string {
	stringRune := []rune(s)
	result := ""
	for i := 0; i <= len(stringRune)-1; i++ {
		if (stringRune[i] >= 'a') && (stringRune[i] <= 'z') {
			stringRune[i] = stringRune[i] - 32
		}
		result += string(stringRune[i])
	}
	return result
}

// changes lement to lowercase
func ToLower(s string) string {
	stringRune := []rune(s)
	result := ""
	for i := 0; i <= len(stringRune)-1; i++ {
		if (stringRune[i] >= 'A') && (stringRune[i] <= 'Z') {
			stringRune[i] = stringRune[i] + 32
		}
		result += string(stringRune[i])
	}
	return result
}

// Function to Convert binary number into decimal number

func ConvertBinaryToDecimal(number string) string {
	result, err := strconv.ParseInt(number, 2, 64)
	if err != nil {
		return number
	}
	return strconv.Itoa(int(result))
}

// Function to Convert Hexadecimal to decimal

func ConvertHexToDecimal(hexaString string) string {
	result, err := strconv.ParseInt(hexaString, 16, 64)
	if err != nil {
		return hexaString
	}
	return strconv.Itoa(int(result))
}
