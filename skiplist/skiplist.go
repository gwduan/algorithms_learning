package skiplist

import (
	"errors"
	"math/rand"
)

const MAX_LEVEL = 16

var (
	ErrDuplicate = errors.New("Duplicate")
	ErrNotFound  = errors.New("Not Found")
)

type Element struct {
	value    int
	level    int
	forwards []*Element
}

type SkipList struct {
	head   *Element
	level  int
	length int
}

func newElement(value int, level int) *Element {
	return &Element{
		value:    value,
		level:    level,
		forwards: make([]*Element, level, level),
	}
}

func (e *Element) Value() int {
	return e.value
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   newElement(0, MAX_LEVEL),
		level:  0,
		length: 0,
	}
}

func (s *SkipList) Level() int {
	return s.level
}

func (s *SkipList) Length() int {
	return s.length
}

func (s *SkipList) Insert(value int) error {
	newLevel := randomLevel()
	if newLevel > s.level {
		s.level++
		newLevel = s.level
	}
	newNode := newElement(value, newLevel)

	preNodes := make([]*Element, newLevel, newLevel)
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && pre.forwards[i].value < value {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && pre.forwards[i].value == value {
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

func (s *SkipList) Delete(value int) error {
	found := false
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && pre.forwards[i].value < value {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && pre.forwards[i].value == value {
			found = true
			pre.forwards[i] = pre.forwards[i].forwards[i]
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

func (s *SkipList) Find(value int) (*Element, error) {
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && pre.forwards[i].value < value {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && pre.forwards[i].value == value {
			return pre.forwards[i], nil
		}
	}

	return nil, ErrNotFound
}

func (s *SkipList) FindAll() ([]*Element, error) {
	if s.length == 0 {
		return nil, ErrNotFound
	}

	results := make([]*Element, 0, s.length)
	for pre := s.head; pre.forwards[0] != nil; pre = pre.forwards[0] {
		results = append(results, pre.forwards[0])
	}

	return results, nil
}

func (s *SkipList) FindBetween(begin int, end int) ([]*Element, error) {
	if s.length == 0 {
		return nil, ErrNotFound
	}

	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && pre.forwards[i].value < begin {
			pre = pre.forwards[i]
		}
	}

	results := make([]*Element, 0, 100)
	for pre.forwards[0] != nil && pre.forwards[0].value < end {
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
