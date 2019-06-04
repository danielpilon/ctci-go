package e04

import (
	"unicode"
)

func isPermutationOfPalindrome(str string) bool {
	var freq [26]int
	odds := 0
	a := int('a')
	for _, c := range str {
		if !unicode.IsLetter(c) {
			continue
		}
		r := int(unicode.ToLower(c)) - a

		freq[r]++

		if freq[r]%2 == 1 {
			odds++
		} else {
			odds--
		}
	}
	return odds < 2
}
