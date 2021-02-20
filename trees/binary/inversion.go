package binary

// Invert inverts a binary tree.
func Invert(t *Tree) *Tree {
	if t == nil {
		return t
	}

	stack := []*Tree{t}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		node.Left, node.Right = node.Right, node.Left
	}

	return t
}
