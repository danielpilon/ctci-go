package math

import "testing"

type testCase struct {
	input    int
	expected int
}

func TestAbs(t *testing.T) {
	cases := []testCase{
		{-1, 1},
		{1, 1},
		{0, 0},
	}

	for _, c := range cases {
		actual := Abs(c.input)
		if actual != c.expected {
			t.Fatalf("got %v for %v, expected %v", actual, c.input, c.expected)
		}
	}
}
