package e05

import "testing"

type testCase struct {
	first    string
	second   string
	expected bool
}

func TestIsOneAway(t *testing.T) {
	cases := []testCase{
		{"pale", "pale", true},
		{"pale", "ple", true},
		{"pale", "p13", false},
		{"pales", "pale", true},
		{"pales", "pa13", false},
		{"pale", "bale", true},
		{"pale", "bA1e", false},
		{"pale", "bae", false},
		{"palessssss", "bae", false},
		{"pales", "palesssss", false},
	}

	for _, c := range cases {
		actual := isOneAway(c.first, c.second)
		if actual != c.expected {
			t.Fatalf("for %v and %v, expected %v, got %v", c.first, c.second, c.expected, actual)
		}
	}
}
