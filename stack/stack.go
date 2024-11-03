package stack

import "errors"

var (
	ErrEmpty = errors.New("Stack Empty")
	ErrFull  = errors.New("Stack Full")
)

type Stack interface {
	Push(v any) (any, error)
	Pop() (any, error)
	Top() (any, error)
	Size() int
	IsEmpty() bool
}
