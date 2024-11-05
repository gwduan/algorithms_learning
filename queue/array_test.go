package queue

import "testing"

func TestNewArrayCircularQueue(t *testing.T) {
	q := NewArrayCircularQueue(2)
	newQueueTest(q, t)
}

func TestArrayCircularQueuePutOne(t *testing.T) {
	q := NewArrayCircularQueue(2)
	putOneTest(q, t)
}

func TestArrayCircularQueuePutTwo(t *testing.T) {
	q := NewArrayCircularQueue(2)
	putTwoTest(q, t)
}

func TestArrayCircularQueueLock(t *testing.T) {
	q := NewArrayCircularQueue(1000)
	lockTest(q, t)
}

func TestArrayCircularQueuePutFull(t *testing.T) {
	q := NewArrayCircularQueue(2)
	putFullTest(q, t)
}

func putFullTest(q Queue, t *testing.T) {
	if got, err := q.Put(1); got != 1 || err != nil {
		t.Errorf("Put(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.Put(2); got != 2 || err != nil {
		t.Errorf("Put(2) == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if _, err := q.Put(3); err != ErrFull {
		t.Errorf("Put(3) should fail because queue is full")
	}

	if got, err := q.Get(); got != 1 || err != nil {
		t.Errorf("Get() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.Put(3); got != 3 || err != nil {
		t.Errorf("Put(3) == (%v, %v), want (%v, %v)", got, err, 3, nil)
	}
	if got := q.Size(); got != 2 {
		t.Errorf("Size() == %v, want %v", got, 2)
	}
	if got := q.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if got, err := q.Head(); got != 2 || err != nil {
		t.Errorf("Head() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if got, err := q.Tail(); got != 3 || err != nil {
		t.Errorf("Tail() == (%v, %v), want (%v, %v)", got, err, 3, nil)
	}

	if got, err := q.Get(); got != 2 || err != nil {
		t.Errorf("Get() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if got, err := q.Put(4); got != 4 || err != nil {
		t.Errorf("Put(4) == (%v, %v), want (%v, %v)", got, err, 4, nil)
	}
	if got := q.Size(); got != 2 {
		t.Errorf("Size() == %v, want %v", got, 2)
	}
	if got := q.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if got, err := q.Head(); got != 3 || err != nil {
		t.Errorf("Head() == (%v, %v), want (%v, %v)", got, err, 3, nil)
	}
	if got, err := q.Tail(); got != 4 || err != nil {
		t.Errorf("Tail() == (%v, %v), want (%v, %v)", got, err, 4, nil)
	}
}
