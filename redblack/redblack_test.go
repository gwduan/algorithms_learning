package redblack

import (
	"errors"
	"testing"
)

func cmpInts(a, b any) int {
	return a.(int) - b.(int)
}

type orderFunc func(*Element, []int) []int

func TestNewRedBlackTree(t *testing.T) {
	r := NewRedBlackTree(cmpInts)

	if got := r.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}
	if got, err := r.Find(1); got != nil || err != ErrNotFound {
		t.Errorf("Find(1) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
}

func TestRedBlackTreeInsert(t *testing.T) {
	r := NewRedBlackTree(cmpInts)
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
	/* The tree shape:
	                   5
	                 /   \
	                /     \
	               /       \
	              /         \
	             /           \
	            2             7
	          /   \         /   \
	         /     \       /     \
	        /       \     /       \
	       1         4   6         9
	      /         /             /
	     /         /             /
	    /         /             /
	   0r        3r            8r
	*/
	if got := r.Length(); got != 10 {
		t.Errorf("Length() = %v, want %v", got, 10)
	}
	if got, err := r.Find(0); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(0) = (%v,%v), want (%vr,%v)", got, err, 0, nil)
	}
	if got, err := r.Find(1); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(1) = (%v,%v), want (%vb,%v)", got, err, 1, nil)
	}
	if got, err := r.Find(2); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(2) = (%v,%v), want (%vb,%v)", got, err, 2, nil)
	}
	if got, err := r.Find(3); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(3) = (%v,%v), want (%vr,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(4); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(4) = (%v,%v), want (%vb,%v)", got, err, 4, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(5) = (%v,%v), want (%vb,%v)", got, err, 5, nil)
	}
	if got, err := r.Find(6); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(6) = (%v,%v), want (%vb,%v)", got, err, 6, nil)
	}
	if got, err := r.Find(7); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(7) = (%v,%v), want (%vb,%v)", got, err, 7, nil)
	}
	if got, err := r.Find(8); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(8) = (%v,%v), want (%vr,%v)", got, err, 8, nil)
	}
	if got, err := r.Find(9); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(9) = (%v,%v), want (%vb,%v)", got, err, 9, nil)
	}
	if got, err := r.Find(10); got != nil || err != ErrNotFound {
		t.Errorf("Find(10) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}

	r.Insert(3)
	r.Insert(5)
	if got, err := r.Find(3); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(3) = (%v,%v), want (%vr,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(5) = (%v,%v), want (%vb,%v)", got, err, 5, nil)
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

func color(color bool) string {
	if color {
		return "r"
	}

	return "b"
}

func TestRedBlackTreeDeleteMax(t *testing.T) {
	r := NewRedBlackTree(cmpInts)
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

	r.DeleteMax()
	/* The tree shape after delete:
	                   5
	                 /   \
	                /     \
	               /       \
	              /         \
	             /           \
	            2             7
	          /   \         /   \
	         /     \       /     \
	        /       \     /       \
	       1         4   6         8
	      /         /
	     /         /
	    /         /
	   0r        3r
	*/

	if got := r.Length(); got != 9 {
		t.Errorf("Length() = %v, want %v", got, 9)
	}
	if got, err := r.Find(9); got != nil || err != ErrNotFound {
		t.Errorf("Find(9) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	if got, err := r.Find(0); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(0) = (%v,%v), want (%vr,%v)", got, err, 0, nil)
	}
	if got, err := r.Find(1); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(1) = (%v,%v), want (%vb,%v)", got, err, 1, nil)
	}
	if got, err := r.Find(2); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(2) = (%v,%v), want (%vb,%v)", got, err, 2, nil)
	}
	if got, err := r.Find(3); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(3) = (%v,%v), want (%vr,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(4); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(4) = (%v,%v), want (%vb,%v)", got, err, 4, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(5) = (%v,%v), want (%vb,%v)", got, err, 5, nil)
	}
	if got, err := r.Find(6); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(6) = (%v,%v), want (%vb,%v)", got, err, 6, nil)
	}
	if got, err := r.Find(7); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(7) = (%v,%v), want (%vb,%v)", got, err, 7, nil)
	}
	if got, err := r.Find(8); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(8) = (%v,%v), want (%vb,%v)", got, err, 8, nil)
	}

	gots, _ := inOrderValues(r)
	wants := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(r)
	wants = []int{5, 2, 1, 0, 4, 3, 7, 6, 8}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(r)
	wants = []int{0, 1, 3, 4, 2, 6, 8, 7, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestRedBlackTreeDeleteMin(t *testing.T) {
	r := NewRedBlackTree(cmpInts)
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

	r.DeleteMin()
	/* The tree shape after delete:
	               5
	             /   \
	            /     \
	           /       \
	          /         \
	         /           \
	        2             7
	      /   \         /   \
	     /     \       /     \
	    /       \     /       \
	   1         4   6         9
	            /             /
	           /             /
	          /             /
	         3r            8r
	*/

	if got := r.Length(); got != 9 {
		t.Errorf("Length() = %v, want %v", got, 9)
	}
	if got, err := r.Find(0); got != nil || err != ErrNotFound {
		t.Errorf("Find(0) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	if got, err := r.Find(1); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(1) = (%v,%v), want (%vb,%v)", got, err, 1, nil)
	}
	if got, err := r.Find(2); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(2) = (%v,%v), want (%vb,%v)", got, err, 2, nil)
	}
	if got, err := r.Find(3); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(3) = (%v,%v), want (%vr,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(4); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(4) = (%v,%v), want (%vb,%v)", got, err, 4, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(5) = (%v,%v), want (%vb,%v)", got, err, 5, nil)
	}
	if got, err := r.Find(6); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(6) = (%v,%v), want (%vb,%v)", got, err, 6, nil)
	}
	if got, err := r.Find(7); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(7) = (%v,%v), want (%vb,%v)", got, err, 7, nil)
	}
	if got, err := r.Find(8); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(8) = (%v,%v), want (%vr,%v)", got, err, 8, nil)
	}
	if got, err := r.Find(9); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(9) = (%v,%v), want (%vb,%v)", got, err, 9, nil)
	}

	gots, _ := inOrderValues(r)
	wants := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(r)
	wants = []int{5, 2, 1, 4, 3, 7, 6, 9, 8}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(r)
	wants = []int{1, 3, 4, 2, 6, 8, 9, 7, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestRedBlackTreeDeleteNonBottomValue(t *testing.T) {
	r := NewRedBlackTree(cmpInts)
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

	r.DeleteValue(2)
	/* The tree shape after delete:
	                   5
	                 /   \
	                /     \
	               /       \
	              /         \
	             /           \
	            3             7
	          /   \         /   \
	         /     \       /     \
	        /       \     /       \
	       1         4   6         9
	      /                       /
	     /                       /
	    /                       /
	   0r                      8r
	*/

	if got := r.Length(); got != 9 {
		t.Errorf("Length() = %v, want %v", got, 9)
	}
	if got, err := r.Find(2); got != nil || err != ErrNotFound {
		t.Errorf("Find(2) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	if got, err := r.Find(0); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(0) = (%v,%v), want (%vr,%v)", got, err, 0, nil)
	}
	if got, err := r.Find(1); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(1) = (%v,%v), want (%vb,%v)", got, err, 1, nil)
	}
	if got, err := r.Find(3); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(3) = (%v,%v), want (%vb,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(4); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(4) = (%v,%v), want (%vb,%v)", got, err, 4, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(5) = (%v,%v), want (%vb,%v)", got, err, 5, nil)
	}
	if got, err := r.Find(6); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(6) = (%v,%v), want (%vb,%v)", got, err, 6, nil)
	}
	if got, err := r.Find(7); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(7) = (%v,%v), want (%vb,%v)", got, err, 7, nil)
	}
	if got, err := r.Find(8); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(8) = (%v,%v), want (%vr,%v)", got, err, 8, nil)
	}
	if got, err := r.Find(9); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(9) = (%v,%v), want (%vb,%v)", got, err, 9, nil)
	}

	gots, _ := inOrderValues(r)
	wants := []int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(r)
	wants = []int{5, 3, 1, 0, 4, 7, 6, 9, 8}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(r)
	wants = []int{0, 1, 4, 3, 6, 8, 9, 7, 5}
	if len(gots) != len(wants) {
		t.Errorf("postOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("postOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}
}

func TestRedBlackTreeDeleteBottomValue(t *testing.T) {
	r := NewRedBlackTree(cmpInts)
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

	r.DeleteValue(6)
	/* The tree shape after delete:
	                   5
	                 /   \
	                /     \
	               /       \
	              /         \
	             /           \
	            2             8
	          /   \         /   \
	         /     \       /     \
	        /       \     /       \
	       1         4   7         9
	      /         /
	     /         /
	    /         /
	   0r        3r
	*/

	if got := r.Length(); got != 9 {
		t.Errorf("Length() = %v, want %v", got, 9)
	}
	if got, err := r.Find(6); got != nil || err != ErrNotFound {
		t.Errorf("Find(6) = (%v,%v), want (%v,%v)", got, err, nil, ErrNotFound)
	}
	if got, err := r.Find(0); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(0) = (%v,%v), want (%vr,%v)", got, err, 0, nil)
	}
	if got, err := r.Find(1); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(1) = (%v,%v), want (%vb,%v)", got, err, 1, nil)
	}
	if got, err := r.Find(2); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(2) = (%v,%v), want (%vb,%v)", got, err, 2, nil)
	}
	if got, err := r.Find(3); got == nil || err != nil || color(got.color) == "b" {
		t.Errorf("Find(3) = (%v,%v), want (%vr,%v)", got, err, 3, nil)
	}
	if got, err := r.Find(4); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(4) = (%v,%v), want (%vb,%v)", got, err, 4, nil)
	}
	if got, err := r.Find(5); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(5) = (%v,%v), want (%vb,%v)", got, err, 5, nil)
	}
	if got, err := r.Find(7); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(7) = (%v,%v), want (%vb,%v)", got, err, 7, nil)
	}
	if got, err := r.Find(8); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(8) = (%v,%v), want (%vb,%v)", got, err, 8, nil)
	}
	if got, err := r.Find(9); got == nil || err != nil || color(got.color) == "r" {
		t.Errorf("Find(9) = (%v,%v), want (%vb,%v)", got, err, 9, nil)
	}

	gots, _ := inOrderValues(r)
	wants := []int{0, 1, 2, 3, 4, 5, 7, 8, 9}
	if len(gots) != len(wants) {
		t.Errorf("inOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("inOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = preOrderValues(r)
	wants = []int{5, 2, 1, 0, 4, 3, 8, 7, 9}
	if len(gots) != len(wants) {
		t.Errorf("preOrder numbers = %v, want %v", len(gots), len(wants))
	}
	for i, v := range gots {
		if v != wants[i] {
			t.Errorf("preOrder() no. %v = %v, want %v", i, v, wants[i])
		}
	}

	gots, _ = postOrderValues(r)
	wants = []int{0, 1, 3, 4, 2, 7, 9, 8, 5}
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
	s = append(s, p.value.(int))
	s = inOrder(p.right, s)

	return s
}

func preOrder(p *Element, s []int) []int {
	if p == nil {
		return s
	}

	s = append(s, p.value.(int))
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
	s = append(s, p.value.(int))

	return s
}
