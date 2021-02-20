package binary

// NewBST creates a new binary search tree (BST) from a slice.
func NewBST(slice []int) *Tree {
	if len(slice) == 0 {
		return nil
	}

	tree := &Tree{Value: slice[0]}
	p := tree

	for _, v := range slice {
		for {
			if v > p.Value {
				if p.Right == nil {
					p.Right = &Tree{Value: v}
					p = tree
					break
				} else {
					p = p.Right
				}
			} else if v < p.Value {
				if p.Left == nil {
					p.Left = &Tree{Value: v}
					p = tree
					break
				} else {
					p = p.Left
				}
			} else {
				break
			}
		}
	}

	return tree
}
