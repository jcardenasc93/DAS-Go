package datastructures

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func ShiftBinaryQueue(s *[]BinaryTreeNode) BinaryTreeNode {
	var val BinaryTreeNode
	if len(*s) > 0 {
		val = (*s)[0]
		(*s) = (*s)[1:]
	}
	return val
}

func UnshiftBinaryQueue(s *[]BinaryTreeNode) BinaryTreeNode {
	var val BinaryTreeNode
	if len(*s) > 1 {
		val = (*s)[len(*s)-1]
		(*s) = (*s)[:len(*s)-1]
	}
	return val
}

var BTree = BinaryTreeNode{
	Value: 20,
	Right: &BinaryTreeNode{
		Value: 50,
		Right: &BinaryTreeNode{
			Value: 100,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryTreeNode{
			Value: 30,
			Right: &BinaryTreeNode{
				Value: 45,
				Right: nil,
				Left:  nil,
			},
			Left: &BinaryTreeNode{
				Value: 29,
				Right: nil,
				Left:  nil,
			},
		},
	},
	Left: &BinaryTreeNode{
		Value: 10,
		Right: &BinaryTreeNode{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryTreeNode{
			Value: 5,
			Right: &BinaryTreeNode{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}

var BTree2 = BinaryTreeNode{
	Value: 20,
	Right: &BinaryTreeNode{
		Value: 50,
		Right: nil,
		Left: &BinaryTreeNode{
			Value: 30,
			Right: &BinaryTreeNode{
				Value: 45,
				Right: &BinaryTreeNode{
					Value: 49,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
			Left: &BinaryTreeNode{
				Value: 29,
				Right: nil,
				Left: &BinaryTreeNode{
					Value: 21,
					Right: nil,
					Left:  nil,
				},
			},
		},
	},
	Left: &BinaryTreeNode{
		Value: 10,
		Right: &BinaryTreeNode{
			Value: 15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryTreeNode{
			Value: 5,
			Right: &BinaryTreeNode{
				Value: 7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}
