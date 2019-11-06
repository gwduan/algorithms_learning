package queue

type ArrayCircularQueue struct {
	pool   []interface{}
	head   int
	tail   int
	length int
}

func NewArrayCircularQueue(size int) *ArrayCircularQueue {
	return &ArrayCircularQueue{pool: make([]interface{}, size+1)}
}

func (q *ArrayCircularQueue) Size() int {
	return q.length
}

func (q *ArrayCircularQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *ArrayCircularQueue) EnQueue(e interface{}) (v interface{}, err error) {
	if (q.tail+1)%len(q.pool) == q.head {
		return nil, ErrFull
	}

	q.pool[q.tail] = e
	q.tail = (q.tail + 1) % len(q.pool)
	q.length++

	return e, nil
}

func (q *ArrayCircularQueue) DeQueue() (e interface{}, err error) {
	if q.length == 0 {
		return nil, ErrEmpty
	}

	e = q.pool[q.head]
	q.head = (q.head + 1) % len(q.pool)
	q.length--

	return e, nil
}

func (q *ArrayCircularQueue) Head() (e interface{}, err error) {
	if q.length == 0 {
		return nil, ErrEmpty
	}

	return q.pool[q.head], nil
}

func (q *ArrayCircularQueue) Tail() (e interface{}, err error) {
	if q.length == 0 {
		return nil, ErrEmpty
	}

	return q.pool[(q.tail+len(q.pool)-1)%len(q.pool)], nil
}
