package linkedList

import "fmt"

type linkedList struct {
	head *node
	tail *node
	size int
}

type node struct {
	next  *node
	value interface{}
}

func New() *linkedList {
	return &linkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l linkedList) newNode(value interface{}, next *node) *node {
	return &node{
		next:  next,
		value: value,
	}
}

func (l *linkedList) initIfEmpty(value interface{}) {
	if l.size == 0 {
		n := l.newNode(value, nil)
		l.head = n
		l.tail = n
	}
}

func (l *linkedList) Append(value interface{}) bool {
	l.initIfEmpty(value)

	if l.size > 0 {
		n := l.newNode(value, nil)
		l.tail.next = n
		l.tail = n
	}

	l.size++

	return true
}

func (l *linkedList) Prepend(value interface{}) bool {
	l.initIfEmpty(value)

	if l.size > 0 {
		n := l.newNode(value, l.head)
		l.head = n
	}

	l.size++

	return true
}

func (l *linkedList) SetAt(position int, value interface{}) bool {
	panic("implement me")
}

func (l linkedList) GetAt(position int) interface{} {
	curItem := l.head
	pos := 1

	for {
		if curItem.next == nil {
			return nil
		}

		if position+1 == pos {
			return curItem.value
		}

		if position+1 > l.size {
			return nil
		}

		pos++
		curItem = curItem.next
	}
}

func (l linkedList) DeleteAt(position int) bool {
	panic("implement me")
}

func (l linkedList) Size() int {
	return l.size
}

func (l linkedList) Print() {
	curItem := l.head
	for {
		fmt.Println(curItem.value)
		if curItem.next == nil {
			break
		}
		curItem = curItem.next
	}
}
