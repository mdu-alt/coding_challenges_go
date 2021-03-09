package doubly

// Node represents a node in a doubly linked list.
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

// New creates a new doubly list from an input slice.
func New(slice []int) *Node {
	var (
		head = &Node{}
		curr = head
	)

	for _, v := range slice {
		curr.Next = &Node{Value: v}
		curr.Next.Prev = curr
		curr = curr.Next
	}

	head = head.Next

	if head != nil {
		head.Prev = nil // remove dummy head
	}

	return head
}
