package e02

import (
	"github.com/danielpilon/ctci-go/common/sort"
	"github.com/thoas/go-funk"
)

func checkPermutationWithSorting(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}

	return sort.SortString(first) == sort.SortString(second)
}

func checkPermutationWithArray(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}

	var chars [256]int

	for _, c := range first {
		ascii := int(c)
		chars[ascii]++
	}

	for _, c := range second {
		ascii := int(c)
		chars[ascii]--
		if chars[ascii] < 0 {
			return false
		}
	}

	sum := funk.Reduce(chars, func(acc, elem int) int { return acc + elem }, 0)
	return sum < 2
}
