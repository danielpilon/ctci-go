package e07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input        [][]int
	expected     [][]int
	description  string
	expectsError bool
}

func TestRotateMatrix(t *testing.T) {
	cases := []testCase{
		testCase{
			input: [][]int{
				{0, 1},
				{2, 3},
			},
			expected: [][]int{
				{2, 0},
				{3, 1},
			},
			description:  "2x2 matrix",
			expectsError: false,
		},
		testCase{
			input: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			expected: [][]int{
				{6, 3, 0},
				{7, 4, 1},
				{8, 5, 2},
			},
			description:  "3x3 matrix",
			expectsError: false,
		},
		testCase{
			input: [][]int{
				{0, 1, 2, 3},
				{4, 5, 6, 7},
				{8, 9, 10, 11},
				{12, 13, 14, 15},
			},
			expected: [][]int{
				{12, 8, 4, 0},
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
			},
			description:  "4x4 matrix",
			expectsError: false,
		},
		testCase{
			input:        [][]int{},
			expected:     [][]int{},
			description:  "0x0 matrix",
			expectsError: true,
		},
		testCase{
			input: [][]int{
				{0, 1, 2},
				{3, 4, 5},
			},
			expected: [][]int{
				{2, 0},
				{3, 1},
			},
			description:  "2x3 matrix",
			expectsError: true,
		},
	}

	for _, c := range cases {
		err := rotateMatrix(c.input)
		if c.expectsError {
			assert.Error(t, err, c.description)
		} else {
			assert.Equal(t, c.expected, c.input, c.description)
		}
	}
}
