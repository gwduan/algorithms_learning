package stack

import "sync"

type element struct {
	value any
	next  *element
}

type ListStack struct {
	head   *element
	length int
	mu     sync.RWMutex
}

func NewListStack() *ListStack {
	return &ListStack{}
}

func (s *ListStack) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}

func (s *ListStack) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length == 0
}

func (s *ListStack) Push(e any) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newElement := &element{
		value: e,
		next:  s.head,
	}
	s.head = newElement
	s.length++

	return e, nil
}

func (s *ListStack) Pop() (e any, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.length == 0 {
		return e, ErrEmpty
	}

	headElement := s.head
	s.head = s.head.next
	s.length--

	headElement.next = nil
	return headElement.value, nil
}

func (s *ListStack) Top() (e any, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.length == 0 {
		return e, ErrEmpty
	}

	return s.head.value, nil
}
