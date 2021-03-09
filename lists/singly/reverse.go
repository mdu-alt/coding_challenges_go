package singly

// Reverse reverses a singly linked list.
func Reverse(head *Node) *Node {
	var prev, next *Node

	for curr := head; curr != nil; curr = next {
		next = curr.Next
		curr.Next = prev
		prev = curr
	}

	return prev
}
