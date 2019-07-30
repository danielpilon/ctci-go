// Write code to remove duplicates from an unsorted linked list.
// FOLLOW UP
// How would you solve this problem if a temporary buffer is not allowed?
package e01

import (
	"github.com/danielpilon/ctci-go/common/linkedlist"
)

func RemoveDuplicatesWithMap(head *linkedlist.LinkedListNode) {
	seen := make(map[int]bool)
	var previous *linkedlist.LinkedListNode

	for n := head; n != nil; n = n.Next {
		if seen[n.Value] {
			previous.Next = n.Next
		} else {
			previous = n
			seen[n.Value] = true
		}
	}
}

func RemoveDuplicatesNoExtraSpace(head *linkedlist.LinkedListNode) {
	for n := head; n != nil; n = n.Next {
		r := n
		for r.Next != nil {
			if n.Value == r.Next.Value {
				r.Next = r.Next.Next
			} else {
				r = r.Next
			}
		}
	}
}
