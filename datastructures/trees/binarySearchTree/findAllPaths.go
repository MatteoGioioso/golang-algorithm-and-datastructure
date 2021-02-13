package binarySearchTree

func FindAllPaths(values []int) [][]int {
	tree := New()
	for _, value := range values {
		tree.Insert(value)
	}
	allPaths := make([][]int, 0)
	find(tree.rootNode, make([]int, 0), &allPaths)

	return allPaths
}

func find(n *node, path []int, allPaths *[][]int) {
	path = append(path, n.value)

	if n.right == nil && n.left == nil {
		*allPaths = append(*allPaths, path)
		return
	}

	if n.left != nil {
		 find(n.left, path, allPaths)
	}

	if n.right != nil {
		find(n.right, path, allPaths)
	}
}
