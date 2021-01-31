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

func (t tree) Remove(value interface{}) {
	panic("implement me")
}

func (t tree) Search(value interface{}) {
	panic("implement me")
}

func (t tree) Size() int {
	return t.size
}

func (t tree) LevelOrderTraversal() {
	q := queue.New()
	q.Enqueuing(t.rootNode)
	currentNode := t.rootNode

	for q.Size() > 0 {
		rawNode, _ := q.Dequeuing()
		n := rawNode.(*node)
		currentNode = n
		fmt.Print(n.value, ", ")
		if currentNode.right != nil {
			q.Enqueuing(currentNode.right)
		}

		if currentNode.left != nil {
			q.Enqueuing(currentNode.left)
		}
	}
}

func (t tree) walk(n *node, q queue.Queuer)  {

	if n.left != nil {
		q.Enqueuing(n.left.value)
		t.walk(n.left, q)
	}

	_, _ = q.Dequeuing()
	q.Print()
	fmt.Println()

	//fmt.Print(dequeuing, ", ")
}
