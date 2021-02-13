package binarySearchTree

func Sum(treeValues []int) int {
	tree := New()
	for _, value := range treeValues {
		tree.Insert(value)
	}

	return sum(tree.rootNode)
}

func sum(n *node) int {
	if n == nil {
		return 0
	}

	left := sum(n.left)
	right := sum(n.right)
	return left + right + n.value
}
