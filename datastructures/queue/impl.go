package queue

import (
	"algo/datastructures/linkedList"
	"errors"
)

type queue struct {
	list linkedList.LinkedLister
}

func New() *queue {
	return &queue{list: linkedList.New()}
}

func (q *queue) Enqueuing(value interface{}) error {
	ok := q.list.Append(value)
	if !ok {
		return errors.New("could not append")
	}

	return nil
}

func (q *queue) Dequeuing() (interface{}, error) {
	cachedHead := q.list.Head()
	ok := q.list.DeleteHead()
	if !ok {
		return nil, errors.New("could not dequeue")
	}

	return cachedHead, nil
}

func (q queue) Peeking() interface{}  {
	return q.list.Head()
}

func (q queue) Size() int {
	return q.list.Size()
}



