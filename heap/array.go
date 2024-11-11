package heap

import (
	"sync"
)

type CmpFunc func(any, any) int

type basicArrayHeap struct {
	pool   []any
	length int
	cmp    CmpFunc
	mu     sync.RWMutex
}

type FixedArrayHeap struct {
	basicArrayHeap
}

type ArrayHeap struct {
	basicArrayHeap
}

func (h *basicArrayHeap) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.length
}

func (h *basicArrayHeap) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.length == 0
}

func (h *basicArrayHeap) Head() (e any, err error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if h.length == 0 {
		return e, ErrEmpty
	}

	return h.pool[0], nil
}

func (h *basicArrayHeap) Delete() (e any, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.length == 0 {
		return e, ErrEmpty
	}

	e = h.pool[0]

	h.length--
	h.pool[0] = h.pool[h.length]

	h.downHeapify(0)

	return e, nil
}

func (h *basicArrayHeap) downHeapify(i int) {
	for {
		maxPos := i
		if i*2+1 <= h.length-1 && h.cmp(h.pool[i], h.pool[i*2+1]) < 0 {
			maxPos = i*2 + 1
		}
		if i*2+2 <= h.length-1 && h.cmp(h.pool[maxPos], h.pool[i*2+2]) < 0 {
			maxPos = i*2 + 2
		}
		if maxPos == i {
			break
		}

		h.pool[i], h.pool[maxPos] = h.pool[maxPos], h.pool[i]

		i = maxPos
	}
}

func (h *basicArrayHeap) upHeapify(i int) {
	for (i-1)/2 >= 0 && h.cmp(h.pool[i], h.pool[(i-1)/2]) > 0 {
		h.pool[i], h.pool[(i-1)/2] = h.pool[(i-1)/2], h.pool[i]
		i = (i - 1) / 2
	}
}

func NewFixedArrayHeap(size int, cmp CmpFunc) *FixedArrayHeap {
	return &FixedArrayHeap{
		basicArrayHeap{
			pool: make([]any, size),
			cmp:  cmp,
		},
	}
}

func (h *FixedArrayHeap) Insert(e any) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.length == len(h.pool) {
		return ErrFull
	}

	h.pool[h.length] = e
	h.length++

	h.upHeapify(h.length - 1)

	return nil
}

func NewArrayHeap(initSize int, cmp CmpFunc) *ArrayHeap {
	return &ArrayHeap{
		basicArrayHeap{
			pool: make([]any, initSize),
			cmp:  cmp,
		},
	}
}

func (h *ArrayHeap) Insert(e any) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.length == len(h.pool) {
		newPool := make([]any, len(h.pool)*2)
		copy(newPool, h.pool)
		h.pool = newPool
	}

	h.pool[h.length] = e
	h.length++

	h.upHeapify(h.length - 1)

	return nil
}
