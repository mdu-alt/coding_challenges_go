package singly

// Node represents a node in a singly linked list.
type Node struct {
	Value interface{}
	Next  *Node
}

// New creates a new singly list from an input slice.
func New(slice []int) *Node {
	var (
		head = &Node{}
		curr = head
	)

	for _, v := range slice {
		curr.Next = &Node{Value: v}
		curr = curr.Next
	}

	head = head.Next

	return head
}
