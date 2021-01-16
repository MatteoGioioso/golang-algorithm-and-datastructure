package linkedList

type LinkedList interface {
	Append(value interface{}) bool
	Prepend(value interface{}) bool
	SetAt(position int, value interface{}) bool
	DeleteAt(position int) bool
	GetAt(position int) interface{}
	Size() int
}
