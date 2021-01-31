package stack

type Stacker interface {
	Push(value interface{}) error
	Pop() (interface{}, error)
	Size() int
}
