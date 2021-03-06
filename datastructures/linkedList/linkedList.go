package linkedList

type LinkedLister interface {
	Append(value interface{}) bool
	Prepend(value interface{}) bool
	Delete(node *Node) bool
	Next(node *Node) *Node
	GetAt(position int) interface{}
	Size() int
	DeleteTail() bool
	DeleteHead() bool
	Tail() interface{}
	Head() interface{}
}
