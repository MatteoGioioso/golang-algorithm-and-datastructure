package linkedList

type LinkedLister interface {
	Append(value interface{}) bool
	Prepend(value interface{}) bool
	SetAt(position int, value interface{}) bool
	Delete(node *Node) bool
	GetAt(position int) interface{}
	Size() int
}
