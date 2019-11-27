package redblack

import (
	"errors"
	"testing"
)

type orderFunc func(*Element, []int) []int

func TestRedBlackTree(t *testing.T) {
	r := NewRedBlackTree()
	if got := r.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}
	if got, err := r.Find(1); got != nil || err != ErrNotFound {
		t.Errorf("Find(1) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	r.Insert(8)
	r.Insert(2)
	r.Insert(0)
	r.Insert(7)
	r.Insert(1)
	r.Insert(3)
	r.Insert(9)
	r.Insert(5)
	r.Insert(6)
	r.Insert(4)
	if got := r.Length(); got != 10 {
		t.Errorf("Length() = %v, want %v", got, 10)
	}
	if got, err := r.Find(8); got == nil || err != nil {
		t.Errorf("Find(8) = (%v,%v), want (%v,%v)", got, err, 8, nil)
	}
	if got, err := r.Find(7); got == nil || err != nil {
		t.Errorf("Find(7) = (%v,%v), want (%v,%v)", got, err, 7, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil {
		t.Errorf("Find(5) = (%v,%v), want (%v,%v)", got, err, 5, nil)
	}
	if got, err := r.Find(1); got == nil || err != nil {
		t.Errorf("Find(1) = (%v,%v), want (%v,%v)", got, err, 1, nil)
	}
	if got, err := r.Find(10); got != nil || err != ErrNotFound {
		t.Errorf("Find(10) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	r.Insert(3)
	r.Insert(5)
	if got, err := r.Find(3); got == nil || err != nil {
		t.Errorf("Find(7) = (%v,%v), want (%v,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil {
		t.Errorf("Find(5) = (%v,%v), want (%v,%v)", got, err, 5, nil)
	}

	gots, _ := inOrderValues(r)
	wants := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(r)
	wants = []int{5, 2, 1, 0, 4, 3, 7, 6, 9, 8}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(r)
	wants = []int{0, 1, 3, 4, 2, 6, 8, 9, 7, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func inOrderValues(r *RedBlackTree) ([]int, error) {
	return getOrderValues(r, inOrder)
}

func preOrderValues(r *RedBlackTree) ([]int, error) {
	return getOrderValues(r, preOrder)
}

func postOrderValues(r *RedBlackTree) ([]int, error) {
	return getOrderValues(r, postOrder)
}

func getOrderValues(r *RedBlackTree, of orderFunc) ([]int, error) {
	if r.root == nil {
		return nil, errors.New("empty")
	}

	results := make([]int, 0, r.length)
	results = of(r.root, results)

	return results, nil
}

func inOrder(p *Element, s []int) []int {
	if p == nil {
		return s
	}

	s = inOrder(p.left, s)
	s = append(s, p.value)
	s = inOrder(p.right, s)

	return s
}

func preOrder(p *Element, s []int) []int {
	if p == nil {
		return s
	}

	s = append(s, p.value)
	s = preOrder(p.left, s)
	s = preOrder(p.right, s)

	return s
}

func postOrder(p *Element, s []int) []int {
	if p == nil {
		return s
	}

	s = postOrder(p.left, s)
	s = postOrder(p.right, s)
	s = append(s, p.value)

	return s
}
