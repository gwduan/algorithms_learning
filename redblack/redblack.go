package redblack

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

const (
	BLACK = 0
	RED   = 1
)

type Element struct {
	value int
	left  *Element
	right *Element
	color int
}

type RedBlackTree struct {
	root   *Element
	length int
}

func newElement(value int, color int) *Element {
	return &Element{value: value, color: color}
}

func (e *Element) Value() int {
	return e.value
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (r *RedBlackTree) Length() int {
	return r.length
}

func (r *RedBlackTree) Find(value int) (*Element, error) {
	for p := r.root; p != nil; {
		switch {
		case value < p.value:
			p = p.left
		case value > p.value:
			p = p.right
		default:
			return p, nil
		}
	}

	return nil, ErrNotFound
}

func (r *RedBlackTree) Insert(value int) {
	r.root = r.insertValue(r.root, value)
	r.root.color = BLACK
}

func (r *RedBlackTree) insertValue(e *Element, value int) *Element {
	if e == nil {
		r.length++
		return newElement(value, RED)
	}

	if value < e.value {
		e.left = r.insertValue(e.left, value)
	} else if value > e.value {
		e.right = r.insertValue(e.right, value)
	} else {
		return e
	}

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
	e.color = RED
	e.left.color = BLACK
	e.right.color = BLACK
}
