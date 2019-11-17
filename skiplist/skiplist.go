package skiplist

import "math/rand"

const MAX_LEVEL = 16

type element struct {
	value    int
	level    int
	forwards []*element
}

type SkipList struct {
	head   *element
	level  int
	length int
}

func newElement(value int, level int) *element {
	return &element{
		value:    value,
		level:    level,
		forwards: make([]*element, level, level),
	}
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   newElement(0, MAX_LEVEL),
		level:  0,
		length: 0,
	}
}

func (s *SkipList) Insert(value int) {
	newLevel := randomLevel()
	if newLevel > s.level {
		s.level++
		newLevel = s.level
	}

	newNode := newElement(value, newLevel)

	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && pre.forwards[i].value < value {
			pre = pre.forwards[i]
		}
		if newLevel > i {
			newNode.forwards[i] = pre.forwards[i]
			pre.forwards[i] = newNode
		}
	}

	s.length++
}

func (s *SkipList) Find(value int) *element {
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && pre.forwards[i].value < value {
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && pre.forwards[i].value == value {
			return pre.forwards[i]
		}
	}

	return nil
}

func (s *SkipList) Delete(value int) {
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
	}
}

func randomLevel() int {
	level := 1
	for rand.Float64() < 0.5 && level < MAX_LEVEL {
		level++
	}

	return level
}
