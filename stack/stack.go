package stack

import "errors"

var ErrEmpty = errors.New("stack is empty")

type Stack interface {
	Push(v interface{}) interface{}
	Pop() (interface{}, error)
	Top() (interface{}, error)
	Size() int
	IsEmpty() bool
}
