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
