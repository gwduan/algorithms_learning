package heap

import (
	"testing"
)

func cmpInts(a, b any) int {
	return a.(int) - b.(int)
}

func TestNewFixedArrayHeap(t *testing.T) {
	h := NewFixedArrayHeap(2, cmpInts)
	newHeapTest(h, t)
}

func TestNewArrayHeap(t *testing.T) {
	h := NewArrayHeap(2, cmpInts)
	newHeapTest(h, t)
}

func TestFixedArrayHeapInsert(t *testing.T) {
	h := NewFixedArrayHeap(2, cmpInts)
	if err := h.Insert(1); err != nil {
		t.Errorf("Insert(1) == %v, want %v", err, nil)
	}
	if err := h.Insert(2); err != nil {
		t.Errorf("Insert(1) == %v, want %v", err, nil)
	}
	if err := h.Insert(3); err == nil {
		t.Errorf("Insert(1) == %v, want %v", err, ErrFull)
	}
}

func TestArrayHeapInsert(t *testing.T) {
	h := NewArrayHeap(2, cmpInts)
	for i := 1; i < 10; i++ {
		if err := h.Insert(i); err != nil {
			t.Errorf("Insert(%d) == %v, want %v", i, err, nil)
		}
	}

	if got := h.Size(); got != 9 {
		t.Errorf("Size() == %v, want %v", got, 9)
	}

	wants := []int{9, 8, 6, 7, 3, 2, 5, 1, 4}
	for i, v := range wants {
		if h.pool[i] != v {
			t.Errorf("Heap.pool[%d] == %v, want %v", i, h.pool[i], v)
		}
	}
}

func TestFixedArrayHeapDelete(t *testing.T) {
	h := NewFixedArrayHeap(100, cmpInts)
	heapDeleteTest(h, t)
}

func TestArrayHeapDelete(t *testing.T) {
	h := NewArrayHeap(2, cmpInts)
	heapDeleteTest(h, t)
}
