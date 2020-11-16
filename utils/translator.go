package utils

import (
	"fmt"
)

// TranslateWord - translates a word to gopher language
func TranslateWord(word string) (string, error) {
	// If no word is supplied, return an error
	if len(word) < 1 {
		return "", fmt.Errorf("no word specified")
	}

	// The word starts with a vowel letter
	if isLetterVowel(rune(word[0])) {
		return "g" + word, nil
	}

	// The word starts with "xr"
	if len(word) > 2 && word[:2] == "xr" {
		return "ge" + word, nil
	}

	// The word starts with a consonant letter, followed by "qu"
	if len(word) > 3 && isLetterConsonant(rune(word[0])) && word[1:3] == "qu" {
		return word[3:] + word[:3] + "ogo", nil
	}

	// The word starts with a consonant sound
	if isLetterConsonant(rune(word[0])) && isLetterConsonant(rune(word[1])) {
		if len(word) == 2 {
			return string(word[1]) + string(word[0]) + "ogo", nil
		}

		index := findFirstVowelIndex(word)

		return word[index:] + word[:index] + "ogo", nil
	}

	return word, nil
}

// isLetterVowel - checks to see if a specified rune is a vowel letter
func isLetterVowel(letter rune) bool {
	return letter == 'a' || letter == 'A' ||
		letter == 'e' || letter == 'E' ||
		letter == 'i' || letter == 'I' ||
		letter == 'o' || letter == 'O' ||
		letter == 'u' || letter == 'U'
}

// isLetterConsonant - checks to see if a specified rune is a consonant letter
func isLetterConsonant(letter rune) bool {
	return !isLetterVowel(letter)
}

// findFirstVowelIndex - returns the index of the first vowel letter in a word
//
// Returns -1 if no vowel is found
func findFirstVowelIndex(word string) int {
	for i, c := range word {
		if isLetterVowel(c) {
			return i
		}
	}

	return -1
}
