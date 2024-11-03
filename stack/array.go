package stack

import (
	"sync"
)

type basicArrayStack struct {
	pool   []any
	length int
	mu     sync.RWMutex
}

type FixedArrayStack struct {
	basicArrayStack
}

type ArrayStack struct {
	basicArrayStack
}

func (s *basicArrayStack) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}

func (s *basicArrayStack) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length == 0
}

func (s *basicArrayStack) Pop() (e any, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.length == 0 {
		return e, ErrEmpty
	}

	s.length--
	return s.pool[s.length], nil
}

func (s *basicArrayStack) Top() (e any, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.length == 0 {
		return e, ErrEmpty
	}

	return s.pool[s.length-1], nil
}

func NewFixedArrayStack(size int) *FixedArrayStack {
	return &FixedArrayStack{
		basicArrayStack{
			pool: make([]any, size),
		},
	}
}

func (s *FixedArrayStack) Push(e any) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.length == len(s.pool) {
		return e, ErrFull
	}

	s.pool[s.length] = e
	s.length++

	return e, nil
}

func NewArrayStack(initSize int) *ArrayStack {
	return &ArrayStack{
		basicArrayStack{
			pool: make([]any, initSize),
		},
	}
}

func (s *ArrayStack) Push(e any) (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.length == len(s.pool) {
		newPool := make([]any, len(s.pool)*2)
		copy(newPool, s.pool)
		s.pool = newPool
	}

	s.pool[s.length] = e
	s.length++

	return e, nil
}
