package e01

import (
	"testing"

	ll "github.com/danielpilon/ctci-go/common/linkedlist"
)

type testCase struct {
	input    *ll.LinkedListNode
	expected *ll.LinkedListNode
}

func TestRemoveDuplicatesWithMap(t *testing.T) {
	cases := []struct {
		input    *ll.LinkedListNode
		expected *ll.LinkedListNode
	}{
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    &ll.LinkedListNode{1, &ll.LinkedListNode{2, &ll.LinkedListNode{1, nil}}},
			expected: &ll.LinkedListNode{1, &ll.LinkedListNode{2, nil}},
		},
		{
			input:    &ll.LinkedListNode{1, &ll.LinkedListNode{2, &ll.LinkedListNode{1, &ll.LinkedListNode{3, &ll.LinkedListNode{1, nil}}}}},
			expected: &ll.LinkedListNode{1, &ll.LinkedListNode{2, &ll.LinkedListNode{3, nil}}},
		},
	}

	for _, c := range cases {
		RemoveDuplicatesWithMap(c.input)
		for c.input != nil || c.expected != nil {
			if c.input == nil || c.expected == nil || c.input.Value != c.expected.Value {
				t.Errorf("actual was %v, expected was %v", c.input, c.expected)
				break
			}

			c.input = c.input.Next
			c.expected = c.expected.Next
		}
	}
}

func TestRemoveDuplicatesNoExtraSpace(t *testing.T) {
	cases := []struct {
		input    *ll.LinkedListNode
		expected *ll.LinkedListNode
	}{
		{
			input:    nil,
			expected: nil,
		},
		{
			input:    &ll.LinkedListNode{1, &ll.LinkedListNode{2, &ll.LinkedListNode{1, nil}}},
			expected: &ll.LinkedListNode{1, &ll.LinkedListNode{2, nil}},
		},
		{
			input:    &ll.LinkedListNode{1, &ll.LinkedListNode{2, &ll.LinkedListNode{1, &ll.LinkedListNode{3, &ll.LinkedListNode{1, nil}}}}},
			expected: &ll.LinkedListNode{1, &ll.LinkedListNode{2, &ll.LinkedListNode{3, nil}}},
		},
	}

	for _, c := range cases {
		RemoveDuplicatesNoExtraSpace(c.input)
		for c.input != nil || c.expected != nil {
			if c.input == nil || c.expected == nil || c.input.Value != c.expected.Value {
				t.Errorf("actual was %v, expected was %v", c.input, c.expected)
				break
			}

			c.input = c.input.Next
			c.expected = c.expected.Next
		}
	}
}
