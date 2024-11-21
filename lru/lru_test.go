package lru

import (
	"fmt"
	"testing"
)

func hashInts(key any) int {
	return key.(int)
}

func cmpInts(a, b any) int {
	return a.(int) - b.(int)
}

func TestNewLRUHashTable(t *testing.T) {
	h := NewLRUHashTable(2, 4, hashInts, cmpInts)
	if got, err := h.Get(1); err != ErrNotFound {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	if err := h.Remove(1); err != ErrNotFound {
		t.Errorf("Remove(1) = %v, want %v", err, ErrNotFound)
	}
}

func TestLRUHashTableAddOne(t *testing.T) {
	h := NewLRUHashTable(2, 4, hashInts, cmpInts)
	h.Add(1, 10)
	if got, err := h.Get(1); got != 10 || err != nil {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 10, nil)
	}
	if got, err := h.Get(2); err != ErrNotFound {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	h.Add(1, 100)
	if got, err := h.Get(1); got != 100 || err != nil {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 100, nil)
	}
	if err := h.Remove(1); err != nil {
		t.Errorf("Remove(1) = %v, want %v", err, nil)
	}
	if err := h.Remove(2); err != ErrNotFound {
		t.Errorf("Remove(2) = %v, want %v", err, ErrNotFound)
	}
	if got, err := h.Get(1); err != ErrNotFound {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
}

func TestLRUHashTableAddTwo(t *testing.T) {
	h := NewLRUHashTable(2, 4, hashInts, cmpInts)
	h.Add(1, 10)
	h.Add(2, 20)
	if got, err := h.Get(1); got != 10 || err != nil {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 10, nil)
	}
	if got, err := h.Get(2); got != 20 || err != nil {
		t.Errorf("Get(2) = (%v, %v), want (%v, %v)", got, err, 20, nil)
	}
	if got, err := h.Get(3); err != ErrNotFound {
		t.Errorf("Get(3) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	h.Add(1, 100)
	h.Add(2, 200)
	if got, err := h.Get(1); got != 100 || err != nil {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 100, nil)
	}
	if got, err := h.Get(2); got != 200 || err != nil {
		t.Errorf("Get(2) = (%v, %v), want (%v, %v)", got, err, 200, nil)
	}

	if err := h.Remove(1); err != nil {
		t.Errorf("Remove(1) = %v, want %v", err, nil)
	}
	if err := h.Remove(2); err != nil {
		t.Errorf("Remove(2) = %v, want %v", err, nil)
	}
	if got, err := h.Get(1); err != ErrNotFound {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	if got, err := h.Get(2); err != ErrNotFound {
		t.Errorf("Get(2) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
}

func TestLRUHashTableAddFour(t *testing.T) {
	h := NewLRUHashTable(2, 4, hashInts, cmpInts)
	h.Add(1, 10)
	h.Add(2, 20)
	h.Add(3, 30)
	h.Add(4, 40)
	if got, err := h.Get(1); got != 10 || err != nil {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 10, nil)
	}
	if got, err := h.Get(2); got != 20 || err != nil {
		t.Errorf("Get(2) = (%v, %v), want (%v, %v)", got, err, 20, nil)
	}
	if got, err := h.Get(3); got != 30 || err != nil {
		t.Errorf("Get(3) = (%v, %v), want (%v, %v)", got, err, 30, nil)
	}
	if got, err := h.Get(4); got != 40 || err != nil {
		t.Errorf("Get(4) = (%v, %v), want (%v, %v)", got, err, 40, nil)
	}
	if got, err := h.Get(5); err != ErrNotFound {
		t.Errorf("Get(5) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	h.Add(1, 100)
	h.Add(2, 200)
	h.Add(3, 300)
	h.Add(4, 400)
	if got, err := h.Get(1); got != 100 || err != nil {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 100, nil)
	}
	if got, err := h.Get(2); got != 200 || err != nil {
		t.Errorf("Get(2) = (%v, %v), want (%v, %v)", got, err, 200, nil)
	}
	if got, err := h.Get(3); got != 300 || err != nil {
		t.Errorf("Get(3) = (%v, %v), want (%v, %v)", got, err, 300, nil)
	}
	if got, err := h.Get(4); got != 400 || err != nil {
		t.Errorf("Get(4) = (%v, %v), want (%v, %v)", got, err, 400, nil)
	}

	if err := h.Remove(1); err != nil {
		t.Errorf("Remove(1) = %v, want %v", err, nil)
	}
	if got, err := h.Get(1); err != ErrNotFound {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	if err := h.Remove(2); err != nil {
		t.Errorf("Remove(2) = %v, want %v", err, nil)
	}
	if got, err := h.Get(2); err != ErrNotFound {
		t.Errorf("Get(2) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	if err := h.Remove(3); err != nil {
		t.Errorf("Remove(3) = %v, want %v", err, nil)
	}
	if err := h.Remove(4); err != nil {
		t.Errorf("Remove(4) = %v, want %v", err, nil)
	}
	if got, err := h.Get(3); err != ErrNotFound {
		t.Errorf("Get(3) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
	if got, err := h.Get(4); err != ErrNotFound {
		t.Errorf("Get(4) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
}

func TestLRU(t *testing.T) {
	h := NewLRUHashTable(2, 4, hashInts, cmpInts)
	h.Add(1, 10)
	h.Add(2, 20)
	h.Add(3, 30)
	h.Add(4, 40) // list: 1 2 3 4
	printAll("Add 1 2 3 4", h)
	h.Add(5, 50) // would remove(1), now list: 2 3 4 5
	printAll("Add 5", h)
	if got, err := h.Get(1); err != ErrNotFound {
		t.Errorf("Get(1) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}

	h.Get(2) // now list: 3 4 5 2
	printAll("Get 2", h)
	h.Add(6, 60) // would remove(3), now list: 4 5 2 6
	printAll("Add 6", h)
	if got, err := h.Get(3); err != ErrNotFound {
		t.Errorf("Get(3) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}

	h.Add(4, 400) // now list: 5 2 6 4
	printAll("Add 4", h)
	h.Add(7, 70) // would remove(5), now list: 2 6 4 7
	printAll("Add 7", h)
	if got, err := h.Get(5); err != ErrNotFound {
		t.Errorf("Get(5) = (%v, %v), want (%v, %v)", got, err, 0, ErrNotFound)
	}
}

func printAll(prefix string, h *LRUHashTable) {
	fmt.Printf("%s: ", prefix)
	for p := h.list; p.next != h.list; p = p.next {
		fmt.Printf("(%v->%v) ", p.next.key, p.next.value)
	}
	fmt.Println()
}
