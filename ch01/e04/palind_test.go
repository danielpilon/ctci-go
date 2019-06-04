package e04

import "testing"

type testCase struct {
	input    string
	expected bool
}

var cases []testCase

func init() {
	cases = []testCase{
		{"Tact Coa", true},
		{"Tact Coat", false},
		{"aa", true},
		{" ", true},
		{"", true},
		{"abcddcbabc", false},
	}
}

func TestIsPermutationOfPalindrome(t *testing.T) {
	for _, c := range cases {
		actual := isPermutationOfPalindrome(c.input)
		if actual != c.expected {
			t.Fatalf("expected %v got %v for %v", c.expected, actual, c.input)
		}
	}
}
