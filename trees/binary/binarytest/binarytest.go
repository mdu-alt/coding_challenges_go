package binarytest

import (
	"github.com/mdu-alt/coding_challenges_go/trees/binary"
)

// TODO
var (
	Empty *binary.Tree = nil

	Simple = &binary.Tree{
		Value: 4,
		Left: &binary.Tree{
			Value: 8,
		},
		Right: &binary.Tree{
			Value: 3,
		},
	}
	SimpleInvert = &binary.Tree{
		Value: 4,
		Left: &binary.Tree{
			Value: 3,
		},
		Right: &binary.Tree{
			Value: 8,
		},
	}

	Left = &binary.Tree{
		Value: 5,
		Left: &binary.Tree{
			Value: 12,
			Left: &binary.Tree{
				Value: 0,
				Left: &binary.Tree{
					Value: 9,
					Left: &binary.Tree{
						Value: -21,
					},
				},
			},
		},
	}
	LeftInvert = &binary.Tree{
		Value: 5,
		Right: &binary.Tree{
			Value: 12,
			Right: &binary.Tree{
				Value: 0,
				Right: &binary.Tree{
					Value: 9,
					Right: &binary.Tree{
						Value: -21,
					},
				},
			},
		},
	}

	Right = &binary.Tree{
		Value: 31,
		Right: &binary.Tree{
			Value: -9,
			Right: &binary.Tree{
				Value: 78,
				Right: &binary.Tree{
					Value: 1,
					Right: &binary.Tree{
						Value: -25,
					},
				},
			},
		},
	}
	RightInvert = &binary.Tree{
		Value: 31,
		Left: &binary.Tree{
			Value: -9,
			Left: &binary.Tree{
				Value: 78,
				Left: &binary.Tree{
					Value: 1,
					Left: &binary.Tree{
						Value: -25,
					},
				},
			},
		},
	}

	Any = &binary.Tree{
		Value: 2,
		Left: &binary.Tree{
			Value: 7,
			Left: &binary.Tree{
				Value: 9,
			},
			Right: &binary.Tree{
				Value: 6,
				Left: &binary.Tree{
					Value: 5,
				},
				Right: &binary.Tree{
					Value: 11,
				},
			},
		},
		Right: &binary.Tree{
			Value: 5,
			Right: &binary.Tree{
				Value: 1,
				Left: &binary.Tree{
					Value: 3,
				},
			},
		},
	}
	AnyInverted = &binary.Tree{
		Value: 2,
		Left: &binary.Tree{
			Value: 5,
			Left: &binary.Tree{
				Value: 1,
				Right: &binary.Tree{
					Value: 3,
				},
			},
		},
		Right: &binary.Tree{
			Value: 7,
			Left: &binary.Tree{
				Value: 6,
				Left: &binary.Tree{
					Value: 11,
				},
				Right: &binary.Tree{
					Value: 5,
				},
			},
			Right: &binary.Tree{
				Value: 9,
			},
		},
	}
)

// New creates a new binary tree from a JSON structure in file s.
// func New(t *testing.T, s string) *binary.Tree {
// 	t.Helper()

// 	jsonFile, err := ioutil.ReadFile(s)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	tree := new(binary.Tree)
// 	if err = json.Unmarshal(jsonFile, tree); err != nil {
// 		t.Fatal(err)
// 	}

// 	return tree
// }
