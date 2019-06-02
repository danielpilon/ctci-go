package sort

import "sort"

type sortRunes []rune

func (r sortRunes) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r sortRunes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r sortRunes) Len() int {
	return len(r)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
