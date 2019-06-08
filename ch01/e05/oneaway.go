// There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away.
// EXAMPLE
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bae -> false
package e05

import mth "github.com/danielpilon/ctci-go/common/math"

func isOneAway(first string, second string) bool {

	fl := len(first)
	sl := len(second)

	if mth.Abs(fl-sl) > 1 {
		return false
	} else if fl == sl {
		return oneReplaceAway(first, second)
	} else if fl == sl+1 {
		return oneInsertAway(first, second)
	} else if fl+1 == sl {
		return oneInsertAway(second, first)
	}

	return false
}

func oneReplaceAway(first string, second string) bool {
	foundReplace := false

	f := []rune(first)
	s := []rune(second)

	for i := range f {
		if f[i] != s[i] {
			if foundReplace {
				return false
			}
			foundReplace = true
		}
	}

	return true
}

func oneInsertAway(first string, second string) bool {
	foundInsert := false

	f := []rune(first)
	s := []rune(second)

	indf := 0
	inds := 0

	for indf < len(f) && inds < len(s) {
		if f[indf] != s[inds] {
			if foundInsert {
				return false
			} else {
				foundInsert = true
				indf++
			}
		} else {
			indf++
			inds++
		}
	}

	return true
}
