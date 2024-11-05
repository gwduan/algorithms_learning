package queue

import "errors"

var (
	ErrEmpty = errors.New("Queue Empty")
	ErrFull  = errors.New("Queue Full")
)

type Queue interface {
	Put(v any) (any, error)
	Get() (any, error)
	Head() (any, error)
	Tail() (any, error)
	Size() int
	IsEmpty() bool
}
