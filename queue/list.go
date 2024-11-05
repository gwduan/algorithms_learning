package queue

import "sync"

type element struct {
	value any
	next  *element
}

type ListQueue struct {
	head   *element
	tail   *element
	length int
	mu     sync.RWMutex
}

func NewListQueue() *ListQueue {
	return &ListQueue{}
}

func (q *ListQueue) Size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.length
}

func (q *ListQueue) IsEmpty() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.length == 0
}

func (q *ListQueue) Put(e any) (any, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	newElement := &element{value: e}
	if q.length == 0 {
		q.tail = newElement
		q.head = newElement
	} else {
		q.tail.next = newElement
		q.tail = q.tail.next
	}
	q.length++

	return e, nil
}

func (q *ListQueue) Get() (e any, err error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.length == 0 {
		return e, ErrEmpty
	}

	headElement := q.head
	q.head = q.head.next
	q.length--
	if q.length == 0 {
		q.tail = q.head
	}

	headElement.next = nil

	return headElement.value, nil
}

func (q *ListQueue) Head() (e any, err error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.length == 0 {
		return e, ErrEmpty
	}

	return q.head.value, nil
}

func (q *ListQueue) Tail() (e any, err error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.length == 0 {
		return e, ErrEmpty
	}

	return q.tail.value, nil
}
