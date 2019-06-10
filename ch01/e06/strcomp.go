// Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).
package e06

import (
	"strconv"
	"strings"
)

func compress(str string) string {

	var comp strings.Builder
	l := len(str)

	count := 0
	for i := 0; i < l; i++ {
		count++
		if i+1 >= l || str[i] != str[i+1] {
			comp.WriteString(string(str[i]))
			comp.WriteString(strconv.Itoa(count))
			count = 0
		}
	}

	if comp.Len() > l {
		return str
	}

	return comp.String()
}
