// Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: if implementing in Java, please use a character array so that you can
// perform this operation in place.)
package e03

func urlify(s []rune, length int) {
	l := len(s)

	if l == length {
		return
	}

	for i := length - 1; i >= 0; i-- {
		c := s[i]
		if c == ' ' {
			s[l-1] = '0'
			s[l-2] = '2'
			s[l-3] = '%'
			l -= 3
		} else {
			s[l-1] = c
			l--
		}
	}
}
