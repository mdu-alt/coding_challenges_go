package binary

// Invert inverts a binary tree.
func Invert(t *Tree) *Tree {
	if t == nil {
		return t
	}

	t.Left, t.Right = t.Right, t.Left

	Invert(t.Left)
	Invert(t.Right)

	return t
}
