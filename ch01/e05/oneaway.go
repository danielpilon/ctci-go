package e05

import mth "github.com/danielpilon/ctci-go/common/math"

func isOneAway(first string, second string) bool {

	fl := len(first)
	sl := len(second)

	if mth.Abs(fl-sl) > 1 {
		return false
	} else if fl == sl {
		return oneReplaceAway(first, second)
	} else if fl > sl {
		return oneInsertAway(first, second)
	} else {
		return oneInsertAway(second, first)
	}
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

	return true
}
