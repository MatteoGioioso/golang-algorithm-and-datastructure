package linkedList

type LinkedLister interface {
	Append(value interface{}) bool
	Prepend(value interface{}) bool
	Delete(node *Node) bool
	GetAt(position int) interface{}
	Size() int
	DeleteTail() bool
	Tail() interface{}
	Head() interface{}
}
