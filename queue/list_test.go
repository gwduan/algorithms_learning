package queue

import "testing"

func TestNewListQueue(t *testing.T) {
	q := NewListQueue()
	newQueueTest(q, t)
}

func TestListQueuePutOne(t *testing.T) {
	q := NewListQueue()
	putOneTest(q, t)
}

func TestListQueuePutTwo(t *testing.T) {
	q := NewListQueue()
	putTwoTest(q, t)
}

func TestListQueueLock(t *testing.T) {
	q := NewListQueue()
	lockTest(q, t)
}
