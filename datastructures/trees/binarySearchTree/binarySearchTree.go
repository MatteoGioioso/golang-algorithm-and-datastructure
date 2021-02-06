package binarySearchTree

type BinarySearchTree interface {
	Insert(value int)
	Remove(value int) error
	Search(value int) (int, error)
}
