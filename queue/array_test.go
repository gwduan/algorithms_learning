package queue

import "testing"

func TestNewArrayCircularQueue(t *testing.T) {
	q := NewArrayCircularQueue(2)
	newQueueTest(q, t)
}

func TestArrayCircularQueueEnQueueOne(t *testing.T) {
	q := NewArrayCircularQueue(2)
	enQueueOneTest(q, t)
}

func TestArrayCircularQueueEnQueueTwo(t *testing.T) {
	q := NewArrayCircularQueue(2)
	enQueueTwoTest(q, t)
}

func TestArrayCircularQueueEnQueueFull(t *testing.T) {
	q := NewArrayCircularQueue(2)
	enQueueFullTest(q, t)
}

func enQueueFullTest(q Queue, t *testing.T) {
	if got, err := q.EnQueue(1); got != 1 || err != nil {
		t.Errorf("EnQueue(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.EnQueue(2); got != 2 || err != nil {
		t.Errorf("EnQueue(2) == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if _, err := q.EnQueue(3); err != ErrFull {
		t.Errorf("EnQueue(3) should fail because queue is full")
	}

	if got, err := q.DeQueue(); got != 1 || err != nil {
		t.Errorf("DeQueue() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.EnQueue(3); got != 3 || err != nil {
		t.Errorf("EnQueue(3) == (%v, %v), want (%v, %v)", got, err, 3, nil)
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

	if got, err := q.DeQueue(); got != 2 || err != nil {
		t.Errorf("DeQueue() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if got, err := q.EnQueue(4); got != 4 || err != nil {
		t.Errorf("EnQueue(4) == (%v, %v), want (%v, %v)", got, err, 4, nil)
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
