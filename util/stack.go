package util

import "errors"

type Stack interface {
	Push(item interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Top() (interface{}, error)
	IsEmpty() bool
	Size() int
}

func NewStack() Stack {
	return &sliceStack{items: make([]interface{}, 0)}
}

type sliceStack struct {
	items []interface{}
}

func (s *sliceStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	return item, nil
}

func (s *sliceStack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *sliceStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *sliceStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	index := len(s.items) - 1
	return s.items[index], nil
}

func (s *sliceStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *sliceStack) Size() int {
	return len(s.items)
}
