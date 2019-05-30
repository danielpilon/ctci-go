package main

import "testing"

var cases []struct {
	input    string
	expected bool
}

func init() {
	cases = []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abcdd", false},
		{"gxSJ2nXDda5WiRq6ow4HNGpxfvS4AzMi2im7UjmfZ3kwGtNf69d69uGns6dL12qeP664hP5sjml85JlxwtA05vRhkK2XnXK3EPPsqHuJmvvkrZByHvYXVV83n7YxcN2h9", false},
		{"", true},
	}
}

func TestIsUniqueWithMap(t *testing.T) {
	runWithCases(t, isUniqueWithMap, "map")
}

func TestIsUniqueWithSlice(t *testing.T) {
	runWithCases(t, isUniqueWithSlice, "slice")
}

func TestIsUniqueWithBitVector(t *testing.T) {
	runWithCases(t, isUniqueWithBitVector, "bit vector")
}

func runWithCases(t *testing.T, isUnique func(string) bool, id string) {
	for _, c := range cases {
		actual := isUnique(c.input)
		if actual != c.expected {
			t.Fatalf("id %v, string %v, expected: %v, actual: %v\n", id, c.input, c.expected, actual)
		}
	}
}
