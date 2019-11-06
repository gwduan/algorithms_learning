package queue

import "testing"

func TestNewArrayCircularQueue(t *testing.T) {
	var q Queue
	q = NewArrayCircularQueue(2)
	if got := q.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := q.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := q.Head(); err != ErrEmpty {
		t.Errorf("Head() should fail because queue is empty")
	}
	if _, err := q.Tail(); err != ErrEmpty {
		t.Errorf("Tail() should fail because queue is empty")
	}
	if _, err := q.DeQueue(); err != ErrEmpty {
		t.Errorf("DeQueue() should fail because queue is empty")
	}
}

func TestEnQueueOne(t *testing.T) {
	var q Queue
	q = NewArrayCircularQueue(2)
	if got, err := q.EnQueue(1); got != 1 || err != nil {
		t.Errorf("EnQueue(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got := q.Size(); got != 1 {
		t.Errorf("Size() == %v, want %v", got, 1)
	}
	if got := q.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if got, err := q.Head(); got != 1 || err != nil {
		t.Errorf("Head() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.Tail(); got != 1 || err != nil {
		t.Errorf("Tail() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}

	if got, err := q.DeQueue(); got != 1 || err != nil {
		t.Errorf("DeQueue() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got := q.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := q.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := q.Head(); err != ErrEmpty {
		t.Errorf("Head() should fail because queue is empty")
	}
	if _, err := q.Tail(); err != ErrEmpty {
		t.Errorf("Tail() should fail because queue is empty")
	}
	if _, err := q.DeQueue(); err != ErrEmpty {
		t.Errorf("DeQueue() should fail because queue is empty")
	}
}

func TestEnQueueTwo(t *testing.T) {
	var q Queue
	q = NewArrayCircularQueue(2)
	if got, err := q.EnQueue(1); got != 1 || err != nil {
		t.Errorf("EnQueue(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.EnQueue(2); got != 2 || err != nil {
		t.Errorf("EnQueue(2) == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if got := q.Size(); got != 2 {
		t.Errorf("Size() == %v, want %v", got, 2)
	}
	if got := q.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if got, err := q.Head(); got != 1 || err != nil {
		t.Errorf("Head() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.Tail(); got != 2 || err != nil {
		t.Errorf("Tail() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}

	if got, err := q.DeQueue(); got != 1 || err != nil {
		t.Errorf("DeQueue() == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got := q.Size(); got != 1 {
		t.Errorf("Size() == %v, want %v", got, 1)
	}
	if got := q.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if got, err := q.Head(); got != 2 || err != nil {
		t.Errorf("Head() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if got, err := q.Tail(); got != 2 || err != nil {
		t.Errorf("Tail() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}

	if got, err := q.DeQueue(); got != 2 || err != nil {
		t.Errorf("DeQueue() == (%v, %v), want (%v, %v)", got, err, 2, nil)
	}
	if got := q.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := q.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := q.Head(); err != ErrEmpty {
		t.Errorf("Head() should fail because queue is empty")
	}
	if _, err := q.Tail(); err != ErrEmpty {
		t.Errorf("Tail() should fail because queue is empty")
	}
	if _, err := q.DeQueue(); err != ErrEmpty {
		t.Errorf("DeQueue() should fail because queue is empty")
	}
}

func TestEnQueueThree(t *testing.T) {
	var q Queue
	q = NewArrayCircularQueue(2)
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
