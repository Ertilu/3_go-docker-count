package count

import (
	"strconv"
	s "strings"
)

func Count(sentence string) (string, string) {
	vowels := []string{"a", "e", "i", "o", "u"}

	// lowercase the user input and then make them an slice
	sLower := s.ToLower(sentence)
	arrInput := s.Split(sLower, "")

	// this var used to notifying is the vowels match ? if match they will change to 1
	// vowelsTotal and consonant is used to notifying how much the consonant alphabet is
	matchA, matchI, matchU, matchE, matchO, vowelsTotal, consonant := 0, 0, 0, 0, 0, 0, 0

	// this var used to notifying for the loops to handle error if no vowels there are
	var noVowels bool = true

	// loops for search is there are a vowels alphabet
	for i := 0; i < len(arrInput); i++ {
		if arrInput[i] == vowels[0] {
			matchA = 1
		} else if arrInput[i] == vowels[1] {
			matchI = 1
		} else if arrInput[i] == vowels[2] {
			matchU = 1
		} else if arrInput[i] == vowels[3] {
			matchE = 1
		} else if arrInput[i] == vowels[4] {
			matchO = 1
		}

		// loops for set the consonant alphabet
		for v := 0; v < len(vowels); v++ {
			if arrInput[i] == vowels[v] {
				vowelsTotal++
				consonant = len(arrInput) - vowelsTotal
				noVowels = false
			}
		}
	}

	// handle error if there are no vowels
	if noVowels {
		consonant = len(arrInput)
	}

	// sum all of vowels alphabet
	vowelsMatch := matchA + matchI + matchU + matchE + matchO

	return "Huruf mati: " + strconv.Itoa(consonant) + ",", "Huruf hidup: " + strconv.Itoa(vowelsMatch)
}
