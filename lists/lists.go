package lists

// NodeSingly represents a node in a singly linked list.
type NodeSingly struct {
	Value interface{}
	Next  *NodeSingly
}

// NodeDoubly represents a node in a doubly linked list.
type NodeDoubly struct {
	Value interface{}
	Prev  *NodeDoubly
	Next  *NodeDoubly
}
