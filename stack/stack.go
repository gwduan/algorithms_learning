package stack

type Stack interface {
	Push(v interface{}) interface{}
	Pop() (interface{}, error)
	Top() (interface{}, error)
	Size() int
	IsEmpty() bool
}
