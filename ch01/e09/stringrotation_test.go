package e09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s1          string
	s2          string
	expected    bool
	description string
}

func TestIsRotation(t *testing.T) {
	cases := []testCase{
		testCase{
			s1:          "waterbottle",
			s2:          "erbottlewat",
			expected:    true,
			description: "erbottlewat should be a rotation of waterbottle",
		},
		testCase{
			s1:          "waterbottle",
			s2:          "waterbottle",
			expected:    true,
			description: "waterbottle should be a rotation of waterbottle",
		},
		testCase{
			s1:          "waterbottleerbottlewate",
			s2:          "erbottlewate",
			expected:    false,
			description: "erbottlewate should not be a rotation of waterbottle",
		},
		testCase{
			s1:          "waterbottle",
			s2:          "",
			expected:    false,
			description: "<empty> should not be a rotation of waterbottle",
		},
		testCase{
			s1:          "",
			s2:          "waterbottle",
			expected:    false,
			description: "waterbottle should not be a rotation of <empty>",
		},
	}

	for _, c := range cases {
		actual := IsRotation(c.s1, c.s2)
		assert.Equal(t, c.expected, actual, c.description)
	}
}
