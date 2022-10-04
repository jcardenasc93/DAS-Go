package datastructures

type BinaryTreeNode struct {
	value int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

var bTree = BinaryTreeNode{
	value: 20,
	right: &BinaryTreeNode{
		value: 50,
		right: &BinaryTreeNode{
			value: 100,
			right: nil,
			left:  nil,
		},
		left: &BinaryTreeNode{
			value: 30,
			right: &BinaryTreeNode{
				value: 45,
				right: nil,
				left:  nil,
			},
			left: &BinaryTreeNode{
				value: 29,
				right: nil,
				left:  nil,
			},
		},
	},
	left: &BinaryTreeNode{
		value: 10,
		right: &BinaryTreeNode{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &BinaryTreeNode{
			value: 5,
			right: &BinaryTreeNode{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}

var bTree2 = BinaryTreeNode{
	value: 20,
	right: &BinaryTreeNode{
		value: 50,
		right: nil,
		left: &BinaryTreeNode{
			value: 30,
			right: &BinaryTreeNode{
				value: 45,
				right: &BinaryTreeNode{
					value: 49,
					left:  nil,
					right: nil,
				},
				left: nil,
			},
			left: &BinaryTreeNode{
				value: 29,
				right: nil,
				left: &BinaryTreeNode{
					value: 21,
					right: nil,
					left:  nil,
				},
			},
		},
	},
	left: &BinaryTreeNode{
		value: 10,
		right: &BinaryTreeNode{
			value: 15,
			right: nil,
			left:  nil,
		},
		left: &BinaryTreeNode{
			value: 5,
			right: &BinaryTreeNode{
				value: 7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}
