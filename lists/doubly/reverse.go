package doubly

// Reverse reverses a doubly linked list.
func Reverse(head *Node) *Node {
	var prev, next *Node

	for curr := head; curr != nil; curr = next {
		next = curr.Next
		curr.Next = prev
		curr.Prev = next
		prev = curr
	}

	return prev
}
