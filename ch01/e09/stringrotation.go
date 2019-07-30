// Assume you have a method isSubstring which checks if one word is a substring
// of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
// call to isSubstring (e.g., "waterbottle" is a rotation of" erbottlewat").
package e09

import "strings"

func IsRotation(s1 string, s2 string) bool {
	s1len := len(s1)
	if s1len == len(s2) && s1len > 0 {
		concat := s1 + s1
		return strings.Contains(concat, s2)
	}
	return false
}
