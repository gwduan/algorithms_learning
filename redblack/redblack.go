package redblack

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
)

const (
	BLACK = false
	RED   = true
)

type Element struct {
	value any
	left  *Element
	right *Element
	color bool
}

type CmpFunc func(any, any) int

type RedBlackTree struct {
	root   *Element
	cmp    CmpFunc
	length int
}

func newElement(value any, color bool) *Element {
	return &Element{
		value: value, color: color,
	}
}

func (e *Element) Value() any {
	return e.value
}

func NewRedBlackTree(cmp CmpFunc) *RedBlackTree {
	return &RedBlackTree{
		cmp: cmp,
	}
}

func (r *RedBlackTree) Length() int {
	return r.length
}

func (r *RedBlackTree) Find(value any) (*Element, error) {
	for p := r.root; p != nil; {
		switch {
		case r.cmp(value, p.value) < 0:
			p = p.left
		case r.cmp(value, p.value) > 0:
			p = p.right
		default:
			return p, nil
		}
	}

	return nil, ErrNotFound
}

func (r *RedBlackTree) Insert(value any) {
	r.root = r.insertValue(r.root, value)
	r.root.color = BLACK
}

func (r *RedBlackTree) insertValue(e *Element, value any) *Element {
	if e == nil {
		r.length++
		return newElement(value, RED)
	}

	if r.cmp(value, e.value) < 0 {
		e.left = r.insertValue(e.left, value)
	} else if r.cmp(value, e.value) > 0 {
		e.right = r.insertValue(e.right, value)
	} else {
		return e
	}

	return fixUp(e)
}

func isRed(e *Element) bool {
	if e == nil {
		return false
	}

	return e.color == RED
}

func rotateLeft(e *Element) *Element {
	x := e.right
	e.right = x.left
	x.left = e
	x.color = e.color
	e.color = RED

	return x
}

func rotateRight(e *Element) *Element {
	x := e.left
	e.left = x.right
	x.right = e
	x.color = e.color
	e.color = RED

	return x
}

func flipColors(e *Element) {
	e.color = !e.color
	e.left.color = !e.left.color
	e.right.color = !e.right.color
}

func fixUp(e *Element) *Element {
	if isRed(e.right) && !isRed(e.left) {
		e = rotateLeft(e)
	}

	if isRed(e.left) && isRed(e.left.left) {
		e = rotateRight(e)
	}

	if isRed(e.left) && isRed(e.right) {
		flipColors(e)
	}

	return e
}

func balance(e *Element) *Element {
	if isRed(e.right) {
		e = rotateLeft(e)
	}

	if isRed(e.left) && isRed(e.left.left) {
		e = rotateRight(e)
	}

	if isRed(e.left) && isRed(e.right) {
		flipColors(e)
	}

	return e
}

func (r *RedBlackTree) DeleteMax() {
	if r.root == nil {
		return
	}

	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = RED
	}

	r.root = r.deleteMax(r.root)
	if r.root != nil {
		r.root.color = BLACK
	}
	r.length--
}

func (r *RedBlackTree) deleteMax(e *Element) *Element {
	if isRed(e.left) {
		e = rotateRight(e)
	}

	if e.right == nil {
		return nil
	}

	if !isRed(e.right) && !isRed(e.right.left) {
		e = moveRedRight(e)
	}

	e.right = r.deleteMax(e.right)

	return balance(e)
}

func (r *RedBlackTree) DeleteMin() {
	if r.root == nil {
		return
	}

	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = RED
	}

	r.root = r.deleteMin(r.root)
	if r.root != nil {
		r.root.color = BLACK
	}
	r.length--
}

func (r *RedBlackTree) deleteMin(e *Element) *Element {
	if e.left == nil {
		return nil
	}

	if !isRed(e.left) && !isRed(e.left.left) {
		e = moveRedLeft(e)
	}

	e.left = r.deleteMin(e.left)

	return balance(e)
}

func (r *RedBlackTree) DeleteValue(value any) {
	if r.root == nil {
		return
	}

	if !isRed(r.root.left) && !isRed(r.root.right) {
		r.root.color = RED
	}

	r.root = r.deleteValue(r.root, value)
	if r.root != nil {
		r.root.color = BLACK
	}
	r.length--
}

func (r *RedBlackTree) deleteValue(e *Element, value any) *Element {
	if r.cmp(value, e.value) < 0 {
		if !isRed(e.left) && !isRed(e.left.left) {
			e = moveRedLeft(e)
		}

		e.left = r.deleteValue(e.left, value)
	} else {
		if isRed(e.left) {
			e = rotateRight(e)
		}

		if r.cmp(value, e.value) == 0 && e.right == nil {
			return nil
		}

		if !isRed(e.right) && !isRed(e.right.left) {
			e = moveRedRight(e)
		}

		if r.cmp(value, e.value) == 0 {
			e.value = min(e.right)
			e.right = r.deleteMin(e.right)
		} else {
			e.right = r.deleteValue(e.right, value)
		}
	}

	return balance(e)
}

func min(e *Element) any {
	for e.left != nil {
		e = e.left
	}

	return e.value
}

func moveRedRight(e *Element) *Element {
	flipColors(e)

	if isRed(e.left.left) {
		e = rotateRight(e)
		flipColors(e)
	}

	return e
}

func moveRedLeft(e *Element) *Element {
	flipColors(e)

	if isRed(e.right.left) {
		e.right = rotateRight(e.right)
		e = rotateLeft(e)
		flipColors(e)
	}

	return e
}
