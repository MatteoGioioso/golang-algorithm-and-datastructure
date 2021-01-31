package queue

type Queuer interface {
	Enqueuing(value interface{}) error
	Dequeuing() (interface{}, error)
	Peeking() interface{}
	Size() int
}
