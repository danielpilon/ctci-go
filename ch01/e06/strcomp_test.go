package e06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input    string
	expected string
}

func TestCompress(t *testing.T) {
	cases := []testCase{
		{"aaaaa", "a5"},
		{"a", "a"},
		{"aa", "a2"},
		{"abc", "abc"},
		{"aaaabbbbc", "a4b4c1"},
		{"", ""},
	}
	for _, c := range cases {
		actual := compress(c.input)
		assert.Equal(t, c.expected, actual, "")
	}
}
