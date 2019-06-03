// Implement an algorithm to determine if a string has all unique characters. What if you
// cannot use additional data structures?
package e01

func isUniqueWithMap(s string) bool {
	if len(s) > 128 {
		return false
	}

	chars := make(map[rune]bool)

	for _, c := range s {
		if _, ok := chars[c]; ok {
			return false
		}
		chars[c] = true
	}

	return true
}

func isUniqueWithSlice(s string) bool {
	if len(s) > 128 {
		return false
	}

	chars := make([]bool, 128, 128)

	for _, c := range s {
		ascii := int(c)
		if chars[ascii] {
			return false
		}
		chars[ascii] = true
	}

	return true
}

func isUniqueWithBitVector(s string) bool {
	if len(s) > 128 {
		return false
	}

	vector := 0
	a := uint('a')

	for _, c := range s {
		ascii := uint(c) - a
		if vector&(1<<ascii) > 0 {
			return false
		}
		vector |= (1 << ascii)
	}

	return true
}
