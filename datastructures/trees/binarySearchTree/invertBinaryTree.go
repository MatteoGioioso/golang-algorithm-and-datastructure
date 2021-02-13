package binarySearchTree

func InvertTree(values []int) []int {
	t := New()
	for _, value := range values {
		t.Insert(value)
	}

	invert(t.rootNode)

	return t.LevelOrderTraversal()
}

func invert(n *node) {
	if n.right == nil && n.left == nil {
		return
	}

	tmpLeft := n.left

	if n.left != nil {
		invert(n.left)
		n.left = n.right
	}

	if n.right != nil {
		invert(n.right)
		n.right = tmpLeft
	}
}
