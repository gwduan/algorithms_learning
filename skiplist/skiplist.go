package skiplist

import (
	"errors"
	"math/rand"
	"sync"
)

const MAX_LEVEL = 16

var (
	ErrDuplicate = errors.New("Duplicate")
	ErrNotFound  = errors.New("Not Found")
)

type CmpFunc func(any, any) int

type Element struct {
	value    any
	level    int
	forwards []*Element
}

type SkipList struct {
	head   *Element
	level  int
	length int
	cmp    CmpFunc
	mu     sync.RWMutex
}

func newElement(value any, level int) *Element {
	return &Element{
		value:    value,
		level:    level,
		forwards: make([]*Element, level, level),
	}
}

func (e *Element) Value() any {
	return e.value
}

func NewSkipList(cmp CmpFunc) *SkipList {
	var zero any

	return &SkipList{
		head: newElement(zero, MAX_LEVEL),
		cmp:  cmp,
	}
}

func (s *SkipList) Level() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.level
}

func (s *SkipList) Length() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}

func (s *SkipList) Insert(value any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	newLevel := randomLevel()
	incLevel := false
	if newLevel > s.level {
		incLevel = true
		s.level++
		newLevel = s.level
	}
	newNode := newElement(value, newLevel)

	preNodes := make([]*Element, newLevel, newLevel)
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) < 0 {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) == 0 {
			if incLevel {
				s.level--
			}
			return ErrDuplicate
		}

		if newLevel > i {
			preNodes[i] = pre
		}
	}

	for i := 0; i < newLevel; i++ {
		newNode.forwards[i] = preNodes[i].forwards[i]
		preNodes[i].forwards[i] = newNode
	}

	s.length++

	return nil
}

func (s *SkipList) Delete(value any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	found := false
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) < 0 {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) == 0 {
			found = true
			e := pre.forwards[i]
			pre.forwards[i] = e.forwards[i]
			e.forwards[i] = nil
			if s.head.forwards[i] == nil {
				s.level--
			}
		}
	}

	if found {
		s.length--
		return nil
	}

	return ErrNotFound
}

func (s *SkipList) Find(value any) (*Element, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) < 0 {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) == 0 {
			return pre.forwards[i], nil
		}
	}

	return nil, ErrNotFound
}

func (s *SkipList) FindAll() ([]*Element, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.length == 0 {
		return nil, ErrNotFound
	}

	results := make([]*Element, 0, s.length)
	for pre := s.head; pre.forwards[0] != nil; pre = pre.forwards[0] {
		results = append(results, pre.forwards[0])
	}

	return results, nil
}

func (s *SkipList) FindBetween(begin any, end any) ([]*Element, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.length == 0 {
		return nil, ErrNotFound
	}

	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, begin) < 0 {
			pre = pre.forwards[i]
		}
	}

	results := make([]*Element, 0, 100)
	for pre.forwards[0] != nil && s.cmp(pre.forwards[0].value, end) < 0 {
		results = append(results, pre.forwards[0])
		pre = pre.forwards[0]
	}

	if len(results) == 0 {
		return nil, ErrNotFound
	}

	return results, nil
}

func randomLevel() int {
	level := 1
	for rand.Float64() < 0.5 && level < MAX_LEVEL {
		level++
	}

	return level
}
