package binarySearchTree

import (
	"algo/datastructures/queue"
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
	if n.value > value {
		if n.right == nil {
			n.right = t.newNode(value)
			t.size++
			return
		} else {
			t.insertRcv(value, n.right)
		}
	}

	if n.value < value {
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
	if input - current == 0 {
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
			if currNode.left == nil {
				break
			}
			currNode = currNode.left
			continue
		}

		if value < currNode.value {
			if currNode.right == nil {
				break
			}
			currNode = currNode.right
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
		if currentNode.right != nil {
			q.Enqueuing(currentNode.right)
		}

		if currentNode.left != nil {
			q.Enqueuing(currentNode.left)
		}
	}

	return arr
}

func (t tree) TreeTraversal() {
	panic("not implemented")
}

func (t tree) Remove(value int) error {
	panic("implement me")
}
