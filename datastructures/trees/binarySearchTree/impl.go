package binarySearchTree

import (
	"algo/datastructures/queue"
	"algo/datastructures/stack"
	"fmt"
)

type tree struct {
	rootNode *node
	size     int
}

func New() *tree {
	return &tree{
		size: 0,
	}
}

type node struct {
	right *node
	left  *node
	value int
}

func (t tree) newNode(value int) *node {
	return &node{
		right: nil,
		left:  nil,
		value: value,
	}
}

func (t *tree) Insert(value int) {
	if t.rootNode == nil {
		t.rootNode = t.newNode(value)
		t.size++
		return
	}

	t.insertRcv(value, t.rootNode)
}

func (t *tree) insertRcv(value int, n *node) {
	if value > n.value {
		if n.right == nil {
			n.right = t.newNode(value)
			t.size++
			return
		} else {
			t.insertRcv(value, n.right)
		}
	}

	if value < n.value  {
		if n.left == nil {
			n.left = t.newNode(value)
			t.size++
			return
		} else {
			t.insertRcv(value, n.left)
		}
	}
}

func (t tree) comparator(input, current int) bool {
	if input-current == 0 {
		return true
	}

	return false
}

func (t tree) Search(value int) (int, error) {
	node := t.search(value, t.rootNode)
	if node == nil {
		return 0, fmt.Errorf("not found")
	}

	return node.value, nil
}

func (t tree) search(value int, n *node) *node {
	var currNode = n
	for {
		if t.comparator(currNode.value, value) {
			return currNode
		}

		if value > currNode.value {
			if currNode.right == nil {
				break
			}
			currNode = currNode.right
			continue
		}

		if value < currNode.value {
			if currNode.left == nil {
				break
			}
			currNode = currNode.left
			continue
		}
	}

	return nil
}

func (t tree) Size() int {
	return t.size
}

// Breadth First Search
func (t tree) LevelOrderTraversal() []int {
	var arr []int
	q := queue.New()
	q.Enqueuing(t.rootNode)
	currentNode := t.rootNode

	for q.Size() > 0 {
		rawNode, _ := q.Dequeuing()
		n := rawNode.(*node)
		currentNode = n
		arr = append(arr, n.value)
		if currentNode.left != nil {
			q.Enqueuing(currentNode.left)
		}

		if currentNode.right != nil {
			q.Enqueuing(currentNode.right)
		}
	}

	return arr
}

func (t tree) TreeTraversal() []int {
	var arr []int
	s := stack.New()
	s.Push(t.rootNode)

	var currentNode *node

	for s.Size() > 0 {
		rawNode, _ := s.Pop()
		currentNode = rawNode.(*node)
		arr = append(arr, currentNode.value)

		if currentNode.right != nil {
			s.Push(currentNode.right)
		}

		if currentNode.left != nil {
			s.Push(currentNode.left)
		}
	}

	return arr
}

func (t *tree) Remove(value int) error {
	t.remove(value, t.rootNode)
	t.size--

	return nil
}

func (t *tree) remove(value int, currentNode *node) {
	var currNode = currentNode
	var parentNode = t.rootNode
	for {
		fmt.Println(value, currNode.value)
		if value == currNode.value {
			if t.isLeaf(currNode) {
				t.unlinkChildNode(parentNode, value)
				break
			}

			if t.hasLeftSubTreeOnly(currNode) {
				parentNode.left = currNode.left
				currNode.left = nil
				break
			}

			if t.hasRightSubTreeOnly(currNode) {
				parentNode.right = currNode.right
				currNode.right = nil
				break
			}

			if t.hasBothSubTrees(currNode) {
				smallestNode := t.findSmallestNodeInTheRightSubTree(currNode)
				currNode.value = smallestNode.value

				// Little hack to make the iteration continue and set
				// the new value to remove as the smallest node in the right sub tree
				value = smallestNode.value
				currNode = currNode.right
			}
		}

		if value > currNode.value {
			if currNode.right == nil {
				break
			}

			parentNode = currNode
			currNode = currNode.right
			continue
		}

		if value < currNode.value {
			if currNode.left == nil {
				break
			}

			parentNode = currNode
			currNode = currNode.left
			continue
		}
	}
}

func (t *tree) unlinkChildNode(parentNode *node, value int)  {
	if parentNode.left != nil && value == parentNode.left.value {
		parentNode.left = nil
	}

	if parentNode.right != nil && value == parentNode.right.value {
		parentNode.right = nil
	}
}

func (t tree) findSmallestNodeInTheRightSubTree(currentNode *node) *node {
	leftSubTree := currentNode.right
	smallestNode := leftSubTree

	for {
		if smallestNode.left == nil {
			break
		}

		smallestNode = leftSubTree.left
	}

	return smallestNode
}

func (t tree) isLeaf(n *node) bool {
	return n.left == nil && n.right == nil
}

func (t tree) hasLeftSubTreeOnly(n *node) bool {
	if n.left != nil && n.right == nil {
		return true
	}

	return false
}

func (t tree) hasRightSubTreeOnly(n *node) bool {
	if n.right != nil && n.left == nil {
		return true
	}

	return false
}

func (t tree) hasBothSubTrees(n *node) bool {
	if n.right != nil && n.left != nil {
		return true
	}

	return false
}
