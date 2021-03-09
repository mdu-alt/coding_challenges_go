package lists

// ReverseSingly reverses a singly linked list.
func ReverseSingly(head *NodeSingly) *NodeSingly {
	var prev, next *NodeSingly

	for curr := head; curr != nil; curr = next {
		next = curr.Next
		curr.Next = prev
		prev = curr
	}

	return prev
}

// ReverseDoubly reverses a doubly linked list.
func ReverseDoubly(head *NodeDoubly) *NodeDoubly {
	var prev, next *NodeDoubly

	for curr := head; curr != nil; curr = next {
		next = curr.Next
		curr.Next = prev
		curr.Prev = next
		prev = curr
	}

	return prev
}
