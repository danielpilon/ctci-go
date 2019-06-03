package e03

import "testing"

type testCase struct {
	input       string
	inputLength int
	expected    string
}

var cases []testCase

func init() {
	cases = []testCase{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"MrJohnSmith", 11, "MrJohnSmith"},
		{" MrJohnSmith  ", 12, "%20MrJohnSmith"},
	}
}

func TestUrlify(t *testing.T) {
	for _, c := range cases {
		r := []rune(c.input)
		urlify(r, c.inputLength)
		actual := string(r)
		if actual != c.expected {
			t.Fatalf("%v is not the expected result for %v, expected %v", actual, c.input, c.expected)
		}
	}
}
