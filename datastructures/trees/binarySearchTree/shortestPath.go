package binarySearchTree

import (
	"algo/datastructures/stack"
	"fmt"
)

func ShortestPath(values []int) int {
	tree := New()
	for _, value := range values {
		tree.Insert(value)
	}

	shortest(tree.rootNode)

	return 0
}

func shortest(n *node) {
	s := stack.New()
	s.Push(n)

	var currentNode *node

	for s.Size() > 0 {
		pNode, _ := s.Pop()
		currentNode = pNode.(*node)

		if currentNode.right != nil {
			s.Push(currentNode.right)
		}

		if currentNode.left != nil {
			s.Push(currentNode.left)
		}

		fmt.Println(currentNode.value)
	}
}
