package heap

import (
	"slices"
	"testing"
)

func newHeapTest(h Heap, t *testing.T) {
	if got := h.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := h.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
}

func heapDeleteTest(h Heap, t *testing.T) {
	for i := 1; i < 10; i++ {
		_ = h.Insert(i)
	}

	if v, err := h.Head(); err != nil {
		t.Errorf("Head() == %v, want %v", err, nil)
	} else if v != 9 {
		t.Errorf("Head() == %v, want %v", v, 9)
	}

	gots := make([]int, 0, h.Size())
	for h.Size() > 0 {
		v, err := h.Delete()
		if err != nil {
			t.Errorf("Delete() %v, want %v", err, nil)
		}
		gots = append(gots, v.(int))
	}

	wants := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	if !slices.Equal(gots, wants) {
		t.Errorf("Delete elements is %v, want %v", gots, wants)
	}

	if _, err := h.Delete(); err != ErrEmpty {
		t.Errorf("Delete() %v, want %v", err, ErrEmpty)
	}
}
