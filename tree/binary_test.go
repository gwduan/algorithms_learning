package tree

import (
	"errors"
	"testing"
)

type orderFunc func(*element, []int) []int

func cmpInts(a, b any) int {
	return a.(int) - b.(int)
}

func TestBinarySearchTree(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	if got := b.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}
	if got, err := b.Find(1); got != nil || err != ErrNotFound {
		t.Errorf("Find(1) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	b.Insert(5)
	b.Insert(3)
	b.Insert(8)
	b.Insert(9)
	b.Insert(7)
	b.Insert(1)
	b.Insert(4)
	b.Insert(6)
	b.Insert(2)
	if got := b.Length(); got != 9 {
		t.Errorf("Length() = %v, want %v", got, 9)
	}
	if got, err := b.Find(7); got == nil || err != nil {
		t.Errorf("Find(7) = (%v,%v), want (%v,%v)", got.value, err, 7, nil)
	}
	if got, err := b.Find(10); got != nil || err != ErrNotFound {
		t.Errorf("Find(10) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	if err := b.Insert(4); err != ErrDuplicate {
		t.Errorf("Insert(4) = %v, want %v", err, ErrDuplicate)
	}

	gots, _ := inOrderValues(b)
	wants := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(b)
	wants = []int{5, 3, 1, 2, 4, 8, 7, 6, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(b)
	wants = []int{2, 1, 4, 3, 6, 7, 9, 8, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestBinarySearchTreeLeftRight(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	if got, err := b.Left(); got != nil || err != ErrNotFound {
		t.Errorf("Left() = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	if got, err := b.Right(); got != nil || err != ErrNotFound {
		t.Errorf("Right() = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	b.Insert(5)
	if got, err := b.Left(); got.value != 5 || err != nil {
		t.Errorf("Left() = (%v,%v), want (%v,%v)", got.value, err, 5, nil)
	}
	if got, err := b.Right(); got.value != 5 || err != nil {
		t.Errorf("Right() = (%v,%v), want (%v,%v)", got.value, err, 5, nil)
	}

	b.Insert(3)
	b.Insert(8)
	b.Insert(9)
	b.Insert(7)
	b.Insert(1)
	b.Insert(4)
	b.Insert(6)
	b.Insert(2)
	if got, err := b.Left(); got.value != 1 || err != nil {
		t.Errorf("Left() = (%v,%v), want (%v,%v)", got.value, err, 1, nil)
	}
	if got, err := b.Right(); got.value != 9 || err != nil {
		t.Errorf("Right() = (%v,%v), want (%v,%v)", got.value, err, 9, nil)
	}
}

func TestBinarySearchTreeFloorCeiling(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	b.Insert(50)
	b.Insert(30)
	b.Insert(80)
	b.Insert(90)
	b.Insert(70)
	b.Insert(10)
	b.Insert(40)
	b.Insert(60)
	b.Insert(20)
	if got, err := b.Floor(40); got.value != 40 || err != nil {
		t.Errorf("Floor(40) = (%v,%v), want (%v,%v)", got.value, err, 40, nil)
	}
	if got, err := b.Floor(45); got.value != 40 || err != nil {
		t.Errorf("Floor(45) = (%v,%v), want (%v,%v)", got.value, err, 40, nil)
	}
	if got, err := b.Floor(55); got.value != 50 || err != nil {
		t.Errorf("Floor(55) = (%v,%v), want (%v,%v)", got.value, err, 50, nil)
	}
	if got, err := b.Floor(5); got != nil || err != ErrNotFound {
		t.Errorf("Floor(5) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	if got, err := b.Ceiling(40); got.value != 40 || err != nil {
		t.Errorf("Ceiling(40) = (%v,%v), want (%v,%v)", got.value, err, 40, nil)
	}
	if got, err := b.Ceiling(45); got.value != 50 || err != nil {
		t.Errorf("Ceiling(45) = (%v,%v), want (%v,%v)", got.value, err, 50, nil)
	}
	if got, err := b.Ceiling(85); got.value != 90 || err != nil {
		t.Errorf("Ceiling(85) = (%v,%v), want (%v,%v)", got.value, err, 90, nil)
	}
	if got, err := b.Ceiling(95); got != nil || err != ErrNotFound {
		t.Errorf("Ceiling(95) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
}

func TestBinarySearchTreeDeleteLeafNode(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	if err := b.Delete(5); err != ErrNotFound {
		t.Errorf("Delete(5) = %v, want %v", err, ErrNotFound)
	}

	b.Insert(5)
	if err := b.Delete(5); err != nil {
		t.Errorf("Delete(5) = %v, want %v", err, nil)
	}
	if got := b.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}

	b.Insert(5)
	b.Insert(3)
	b.Insert(8)
	b.Insert(9)
	b.Insert(7)
	b.Insert(1)
	b.Insert(4)
	b.Insert(6)
	b.Insert(2)

	b.Delete(6)
	if got := b.Length(); got != 8 {
		t.Errorf("Length() = %v, want %v", got, 8)
	}
	if got, err := b.Find(6); got != nil || err != ErrNotFound {
		t.Errorf("Find(6) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	gots, _ := inOrderValues(b)
	wants := []int{1, 2, 3, 4, 5, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(b)
	wants = []int{5, 3, 1, 2, 4, 8, 7, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(b)
	wants = []int{2, 1, 4, 3, 7, 9, 8, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestBinarySearchTreeDeleteOneChildNode(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	b.Insert(5)
	b.Insert(3)
	b.Insert(8)
	b.Insert(9)
	b.Insert(7)
	b.Insert(1)
	b.Insert(4)
	b.Insert(6)
	b.Insert(2)

	b.Delete(1)
	if got := b.Length(); got != 8 {
		t.Errorf("Length() = %v, want %v", got, 8)
	}
	if got, err := b.Find(1); got != nil || err != ErrNotFound {
		t.Errorf("Find(1) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	gots, _ := inOrderValues(b)
	wants := []int{2, 3, 4, 5, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(b)
	wants = []int{5, 3, 2, 4, 8, 7, 6, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(b)
	wants = []int{2, 4, 3, 6, 7, 9, 8, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	b.Delete(7)
	if got := b.Length(); got != 7 {
		t.Errorf("Length() = %v, want %v", got, 7)
	}
	if got, err := b.Find(7); got != nil || err != ErrNotFound {
		t.Errorf("Find(7) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	gots, _ = inOrderValues(b)
	wants = []int{2, 3, 4, 5, 6, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(b)
	wants = []int{5, 3, 2, 4, 8, 6, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(b)
	wants = []int{2, 4, 3, 6, 9, 8, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestBinarySearchTreeDeleteTwoChildNode(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	b.Insert(5)
	b.Insert(3)
	b.Insert(8)
	b.Insert(9)
	b.Insert(7)
	b.Insert(1)
	b.Insert(4)
	b.Insert(6)
	b.Insert(2)

	b.Delete(3)
	if got := b.Length(); got != 8 {
		t.Errorf("Length() = %v, want %v", got, 8)
	}
	if got, err := b.Find(3); got != nil || err != ErrNotFound {
		t.Errorf("Find(3) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	gots, _ := inOrderValues(b)
	wants := []int{1, 2, 4, 5, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(b)
	wants = []int{5, 4, 1, 2, 8, 7, 6, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(b)
	wants = []int{2, 1, 4, 6, 7, 9, 8, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestBinarySearchTreeDeleteRootNode(t *testing.T) {
	b := NewBinarySearchTree(cmpInts)
	b.Insert(5)
	b.Insert(3)
	b.Insert(8)
	b.Insert(9)
	b.Insert(7)
	b.Insert(1)
	b.Insert(4)
	b.Insert(6)
	b.Insert(2)

	b.Delete(5)
	if got := b.Length(); got != 8 {
		t.Errorf("Length() = %v, want %v", got, 8)
	}
	if got, err := b.Find(5); got != nil || err != ErrNotFound {
		t.Errorf("Find(5) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	gots, _ := inOrderValues(b)
	wants := []int{1, 2, 3, 4, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(b)
	wants = []int{6, 3, 1, 2, 4, 8, 7, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(b)
	wants = []int{2, 1, 4, 3, 7, 9, 8, 6}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func inOrderValues(b *BinarySearchTree) ([]int, error) {
	return getOrderValues(b, inOrder)
}

func preOrderValues(b *BinarySearchTree) ([]int, error) {
	return getOrderValues(b, preOrder)
}

func postOrderValues(b *BinarySearchTree) ([]int, error) {
	return getOrderValues(b, postOrder)
}

func getOrderValues(b *BinarySearchTree, of orderFunc) ([]int, error) {
	if b.root == nil {
		return nil, errors.New("empty")
	}

	results := make([]int, 0, b.length)
	results = of(b.root, results)

	return results, nil
}

func inOrder(p *element, s []int) []int {
	if p == nil {
		return s
	}

	s = inOrder(p.left, s)
	s = append(s, p.value.(int))
	s = inOrder(p.right, s)

	return s
}

func preOrder(p *element, s []int) []int {
	if p == nil {
		return s
	}

	s = append(s, p.value.(int))
	s = preOrder(p.left, s)
	s = preOrder(p.right, s)

	return s
}

func postOrder(p *element, s []int) []int {
	if p == nil {
		return s
	}

	s = postOrder(p.left, s)
	s = postOrder(p.right, s)
	s = append(s, p.value.(int))

	return s
}
