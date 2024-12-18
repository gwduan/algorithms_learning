package lru

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
)

type HashFunc func(any) int
type CmpFunc func(any, any) int

type element struct {
	pre   *element
	next  *element
	hnext *element
	key   any
	value any
}

type LRUHashTable struct {
	slots     []element
	slotsSize int
	list      *element
	listSize  int
	capacity  int
	hs        HashFunc
	cmp       CmpFunc
}

func newElement(key any, value any) *element {
	return &element{
		key:   key,
		value: value,
	}
}

func NewLRUHashTable(size int, capacity int, hs HashFunc, cmp CmpFunc) *LRUHashTable {
	list := &element{}
	list.next = list
	list.pre = list

	return &LRUHashTable{
		slots:     make([]element, size, size),
		slotsSize: size,
		list:      list,
		listSize:  0,
		capacity:  capacity,
		hs:        hs,
		cmp:       cmp,
	}
}

func (h *LRUHashTable) Get(key any) (value any, err error) {
	_, node := h.getNode(key)
	if node == nil {
		return value, ErrNotFound
	}

	h.moveToListTail(node)

	return node.value, nil
}

func (h *LRUHashTable) Add(key any, value any) {
	i, node := h.getNode(key)
	if node != nil {
		node.value = value
		h.moveToListTail(node)
		return
	}

	node = newElement(key, value)
	node.hnext = h.slots[i].hnext
	h.slots[i].hnext = node
	h.addToListTail(node)

	h.listSize++

	if h.listSize > h.capacity {
		h.Remove(h.list.next.key)
	}
}

func (h *LRUHashTable) Remove(key any) error {
	for p := &h.slots[h.hash(key)]; p.hnext != nil; p = p.hnext {
		if h.cmp(p.hnext.key, key) == 0 {
			h.removeFromList(p.hnext)
			p.hnext = p.hnext.hnext
			h.listSize--
			return nil
		}
	}

	return ErrNotFound
}

func (h *LRUHashTable) hash(key any) int {
	return h.hs(key) % h.slotsSize
}

func (h *LRUHashTable) getNode(key any) (int, *element) {
	i := h.hash(key)
	for p := &h.slots[i]; p.hnext != nil; p = p.hnext {
		if h.cmp(p.hnext.key, key) == 0 {
			return i, p.hnext
		}
	}

	return i, nil
}

func (h *LRUHashTable) moveToListTail(node *element) {
	h.removeFromList(node)
	h.addToListTail(node)
}

func (h *LRUHashTable) removeFromList(node *element) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (h *LRUHashTable) addToListTail(node *element) {
	node.pre = h.list.pre
	node.next = h.list
	h.list.pre.next = node
	h.list.pre = node
}
