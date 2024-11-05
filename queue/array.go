package queue

import "sync"

type ArrayCircularQueue struct {
	pool   []any
	head   int
	tail   int
	length int
	mu     sync.RWMutex
}

func NewArrayCircularQueue(size int) *ArrayCircularQueue {
	return &ArrayCircularQueue{
		pool: make([]any, size+1),
	}
}

func (q *ArrayCircularQueue) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.length
}

func (q *ArrayCircularQueue) IsEmpty() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.length == 0
}

func (q *ArrayCircularQueue) Put(e any) (any, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if (q.tail+1)%len(q.pool) == q.head {
		return e, ErrFull
	}

	q.pool[q.tail] = e
	q.tail = (q.tail + 1) % len(q.pool)
	q.length++

	return e, nil
}

func (q *ArrayCircularQueue) Get() (e any, err error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.length == 0 {
		return e, ErrEmpty
	}

	e = q.pool[q.head]
	q.head = (q.head + 1) % len(q.pool)
	q.length--

	return e, nil
}

func (q *ArrayCircularQueue) Head() (e any, err error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.length == 0 {
		return e, ErrEmpty
	}

	return q.pool[q.head], nil
}

func (q *ArrayCircularQueue) Tail() (e any, err error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.length == 0 {
		return e, ErrEmpty
	}

	return q.pool[(q.tail+len(q.pool)-1)%len(q.pool)], nil
}
