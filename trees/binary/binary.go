package binary

// Tree represents a binary tree.
type Tree struct {
	Value int   `json:"v"`
	Left  *Tree `json:"l,omitempty"`
	Right *Tree `json:"r,omitempty"`
}
