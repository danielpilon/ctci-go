package sort

import "testing"

func TestSortString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"ABCD", "ABCD"},
		{"abcd", "abcd"},
		{"AaBb", "ABab"},
		{"Aa Bb", " ABab"},
	}

	for _, c := range cases {
		actual := SortString(c.input)
		if actual != c.expected {
			t.Fatalf("string %v, expected: %v, actual: %v\n", c.input, c.expected, actual)
		}

	}
}
