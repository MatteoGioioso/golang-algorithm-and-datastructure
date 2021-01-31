package stack

import (
	"algo/datastructures/linkedList"
	"errors"
)

type stack struct {
	list          linkedList.LinkedLister
}

func New() *stack {
	return &stack{
		list:          linkedList.New(),
	}
}

func (s *stack) Push(value interface{}) error {
	ok := s.list.Prepend(value)
	if !ok {
		return errors.New("unable to push")
	}

	return nil
}

func (s *stack) Pop() (interface{}, error) {
	// We pre cache the head because it will delete after
	cachedHead := s.list.Head()
	ok := s.list.DeleteTail()
	if !ok {
		return nil, errors.New("unable to Pop")
	}

	return cachedHead, nil
}

func (s stack) Size() int {
	return s.list.Size()
}
