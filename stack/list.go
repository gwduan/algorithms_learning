package stack

import "errors"

type element struct {
	next  *element
	value interface{}
}

type ListStack struct {
	head   *element
	length int
}

func NewListStack() *ListStack {
	return &ListStack{}
}

func (s *ListStack) Size() int {
	return s.length
}

func (s *ListStack) IsEmpty() bool {
	return s.length == 0
}

func (s *ListStack) Push(e interface{}) interface{} {
	newElement := &element{
		value: e,
		next:  s.head}
	s.head = newElement
	s.length++

	return e
}

func (s *ListStack) Pop() (e interface{}, err error) {
	if s.length == 0 {
		return nil, errors.New("Stack is empty!")
	}

	headElement := s.head
	s.head = s.head.next
	s.length--
	headElement.next = nil

	return headElement.value, nil
}

func (s *ListStack) Top() (e interface{}, err error) {
	if s.length == 0 {
		return nil, errors.New("Stack is empty!")
	}

	return s.head.value, nil
}
