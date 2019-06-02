package e02

import "testing"

type testCase struct {
	first    string
	second   string
	expected bool
}

var cases []testCase

func init() {
	cases = []testCase{
		{"abcd", "dcba", true},
		{"abcde", "dcba", false},
		{"abcde ", " abcde", true},
		{"abcde  ", " abcde", false},
	}
}

func TestCheckPermutationWithSorting(t *testing.T) {
	runWithCases(t, checkPermutationWithSorting, "with sorting")
}

func TestCheckPermutationWithSArray(t *testing.T) {
	runWithCases(t, checkPermutationWithArray, "with array")
}

func runWithCases(t *testing.T, checkPerm func(string, string) bool, id string) {
	for _, c := range cases {
		actual := checkPerm(c.first, c.second)
		if actual != c.expected {
			t.Fatalf("id %v, first %v, second %v, expected: %v, actual: %v\n", id, c.first, c.second, c.expected, actual)
		}
	}
}
