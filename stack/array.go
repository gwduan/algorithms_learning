package stack

type ArrayStack struct {
	pool   []interface{}
	length int
}

func NewArrayStack(size int) *ArrayStack {
	return &ArrayStack{pool: make([]interface{}, size)}
}

func (s *ArrayStack) Size() int {
	return s.length
}

func (s *ArrayStack) IsEmpty() bool {
	return s.length == 0
}

func (s *ArrayStack) Push(e interface{}) interface{} {
	if s.length == len(s.pool) {
		newPool := make([]interface{}, len(s.pool)*2)
		copy(newPool, s.pool)
		s.pool = newPool
	}

	s.pool[s.length] = e
	s.length++

	return e
}

func (s *ArrayStack) Pop() (e interface{}, err error) {
	if s.length == 0 {
		return nil, ErrEmpty
	}

	s.length--
	return s.pool[s.length], nil
}

func (s *ArrayStack) Top() (e interface{}, err error) {
	if s.length == 0 {
		return nil, ErrEmpty
	}

	return s.pool[s.length-1], nil
}
