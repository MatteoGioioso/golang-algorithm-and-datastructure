package linkedList

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	next  *Node
	Value interface{}
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l LinkedList) newNode(value interface{}, next *Node) *Node {
	return &Node{
		next:  next,
		Value: value,
	}
}

func (l *LinkedList) initIfEmpty(value interface{}) {
	if l.size == 0 {
		n := l.newNode(value, nil)
		l.head = n
		l.tail = n
	}
}

func (l *LinkedList) Append(value interface{}) bool {
	l.initIfEmpty(value)

	if l.size > 0 {
		n := l.newNode(value, nil)
		l.tail.next = n
		l.tail = n
	}

	l.size++

	return true
}

func (l *LinkedList) Prepend(value interface{}) bool {
	l.initIfEmpty(value)

	if l.size > 0 {
		n := l.newNode(value, l.head)
		l.head = n
	}

	l.size++

	return true
}

func (l LinkedList) GetAt(position int) interface{} {
	curItem := l.head
	pos := 1

	for {
		if curItem.next == nil {
			return nil
		}

		if position+1 == pos {
			return curItem.Value
		}

		if position+1 > l.size {
			return nil
		}

		pos++
		curItem = curItem.next
	}
}

func (l *LinkedList) Delete(node *Node) bool {
	var currentNode *Node

	if node == l.head {
		l.head = l.head.next
		node.next = nil
		return true
	}

	for {
		currentNode = l.Next(currentNode)
		if currentNode.next == node {
			currentNode.next = currentNode.next.next
			// If the supplied node is the tail
			if node.next == nil {
				l.tail = currentNode
			}

			node.next = nil

			return true
		}
	}
}

func (l *LinkedList) DeleteTail() bool {
	var secondLastNode *Node

	for {
		secondLastNode = l.Next(secondLastNode)
		if secondLastNode.next == nil {
			l.tail = nil
			l.head = nil
			l.size--
			return true
		}

		if secondLastNode.next == l.tail {
			secondLastNode.next = nil
			l.tail = secondLastNode
			l.size--
			return true
		}
	}
}

func (l *LinkedList) DeleteHead() bool {
	oldHead := l.head

	l.head = l.head.next
	oldHead.next = nil
	l.size--

	return true
}


func (l LinkedList) Size() int {
	return l.size
}

func (l LinkedList) Print() {
	curItem := l.head
	for {
		fmt.Println("Val: ", curItem.Value)
		if curItem.next == nil {
			break
		}
		curItem = curItem.next
	}
}

func (l LinkedList) Next(node *Node) *Node {
	// If initial node is nil return the first node
	if node == nil {
		return l.head
	}

	return node.next
}

func (l LinkedList) Tail() interface{} {
	return l.tail.Value
}

func (l LinkedList) Head() interface{} {
	return l.head.Value
}
