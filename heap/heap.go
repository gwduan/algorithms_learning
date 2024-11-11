package heap

import "errors"

var (
	ErrEmpty = errors.New("Heap Empty")
	ErrFull  = errors.New("Heap Full")
)

type Heap interface {
	Insert(any) error
	Delete() (any, error)
	Head() (any, error)
	Size() int
	IsEmpty() bool
}
