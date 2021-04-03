package doubly

// Node represents a node in a doubly linked list.
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

// New creates a new doubly list from an input slice.
func New(slice []int) *Node {
	head := &Node{}
	curr := head

	for _, v := range slice {
		curr.Next = &Node{Value: v}
		curr.Next.Prev = curr
		curr = curr.Next
	}

	if head.Next != nil {
		head.Next.Prev = nil // remove dummy head
	}

	return head.Next
}
