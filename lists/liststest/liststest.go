package liststest

import "github.com/mdu-alt/coding_challenges_go/lists"

// NewSingly creates a new singly list from an input slice.
func NewSingly(slice []int) *lists.NodeSingly {
	var (
		head = &lists.NodeSingly{}
		curr = head
	)

	for _, v := range slice {
		curr.Next = &lists.NodeSingly{Value: v}
		curr = curr.Next
	}

	head = head.Next

	return head
}

// NewDoubly creates a new doubly list from an input slice.
func NewDoubly(slice []int) *lists.NodeDoubly {
	var (
		head = &lists.NodeDoubly{}
		curr = head
	)

	for _, v := range slice {
		curr.Next = &lists.NodeDoubly{Value: v}
		curr.Next.Prev = curr
		curr = curr.Next
	}

	head = head.Next

	if head != nil {
		head.Prev = nil // remove dummy head
	}

	return head
}
