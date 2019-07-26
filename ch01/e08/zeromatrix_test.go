package e08

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

func TestZeroMatrixWithArray(t *testing.T) {
	runWithCases(t, zeroMatrixWithArray, "zeroMatrixWithArray")
}

func TestZeroMatrixInPlace(t *testing.T) {
	runWithCases(t, zeroMatrixInPlace, "zeroMatrixInPlace")
}

func runWithCases(t *testing.T, zeroMatrix func([][]int) error, id string) {
	cases := []testCase{
		testCase{
			input: [][]int{
				{0, 1, 2},
				{3, 4, 5},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 4, 5},
			},
			description:  "2x3 matrix",
			expectsError: false,
		},
		testCase{
			input: [][]int{
				{1, 1, 2},
				{3, 0, 5},
				{3, 4, 5},
			},
			expected: [][]int{
				{1, 0, 2},
				{0, 0, 0},
				{3, 0, 5},
			},
			description:  "3x3 matrix",
			expectsError: false,
		},
		testCase{
			input:        [][]int{},
			expected:     [][]int{},
			description:  "empty matrix",
			expectsError: true,
		},
	}

	for _, c := range cases {
		err := zeroMatrix(c.input)

		if c.expectsError {
			assert.Error(t, err, "id %v, description %v", id, c.description)
		} else {
			assert.Equal(t, c.expected, c.input, "id %v, description %v", id, c.description)
		}
	}
}
