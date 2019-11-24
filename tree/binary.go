package tree

import "errors"

var (
	ErrDuplicate = errors.New("Duplicate")
	ErrNotFound  = errors.New("not found")
)

type element struct {
	value int
	left  *element
	right *element
}

type BinarySearchTree struct {
	root   *element
	length int
}

func newElement(value int) *element {
	return &element{value: value}
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (t *BinarySearchTree) Find(value int) (*element, error) {
	for p := t.root; p != nil; {
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

func (t *BinarySearchTree) Insert(value int) error {
	node := newElement(value)
	if t.root == nil {
		t.root = node
		t.length++
		return nil
	}

	p := t.root
	for {
		switch {
		case value < p.value:
			if p.left == nil {
				p.left = node
				t.length++
				return nil
			}
			p = p.left
		case value > p.value:
			if p.right == nil {
				p.right = node
				t.length++
				return nil
			}
			p = p.right
		default:
			return ErrDuplicate
		}
	}
}
