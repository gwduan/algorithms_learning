package queue

import "errors"

var (
	ErrFull  = errors.New("queue is full")
	ErrEmpty = errors.New("queue is empty")
)

type Queue interface {
	EnQueue(v interface{}) (interface{}, error)
	DeQueue() (interface{}, error)
	Head() (interface{}, error)
	Tail() (interface{}, error)
	Size() int
	IsEmpty() bool
}
