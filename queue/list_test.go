package queue

import "testing"

func TestNewListQueue(t *testing.T) {
	q := NewListQueue()
	newQueueTest(q, t)
}

func TestListQueueEnQueueOne(t *testing.T) {
	q := NewListQueue()
	enQueueOneTest(q, t)
}

func TestListQueueEnQueueTwo(t *testing.T) {
	q := NewListQueue()
	enQueueTwoTest(q, t)
}
