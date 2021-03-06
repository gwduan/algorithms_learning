package stack

import "testing"

func newStackTest(s Stack, t *testing.T) {
	if got := s.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := s.Top(); err != ErrEmpty {
		t.Errorf("Top() should fail because stack is empty")
	}
	if _, err := s.Pop(); err != ErrEmpty {
		t.Errorf("Pop() should fail because stack is empty")
	}
}

func pushOneTest(s Stack, t *testing.T) {
	if got := s.Push(1); got != 1 {
		t.Errorf("Push(1) == %v, want %v", got, 1)
	}
	if got := s.Size(); got != 1 {
		t.Errorf("Size() == %v, want %v", got, 1)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != 1 || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, 1, nil)
	}

	if v, err := s.Pop(); v != 1 || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, 1, nil)
	}
	if got := s.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := s.Top(); err != ErrEmpty {
		t.Errorf("Top() should fail because stack is empty")
	}
	if _, err := s.Pop(); err != ErrEmpty {
		t.Errorf("Pop() should fail because stack is empty")
	}
}

func pushTwoTest(s Stack, t *testing.T) {
	if got := s.Push(1); got != 1 {
		t.Errorf("Push(1) == %v, want %v", got, 1)
	}
	if got := s.Push(2); got != 2 {
		t.Errorf("Push(2) == %v, want %v", got, 2)
	}
	if got := s.Size(); got != 2 {
		t.Errorf("Size() == %v, want %v", got, 2)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != 2 || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, 2, nil)
	}

	if v, err := s.Pop(); v != 2 || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, 2, nil)
	}
	if got := s.Size(); got != 1 {
		t.Errorf("Size() == %v, want %v", got, 1)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != 1 || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, 1, nil)
	}

	if v, err := s.Pop(); v != 1 || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, 1, nil)
	}
	if got := s.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := s.Top(); err != ErrEmpty {
		t.Errorf("Top() should fail because stack is empty")
	}
	if _, err := s.Pop(); err != ErrEmpty {
		t.Errorf("Pop() should fail because stack is empty")
	}
}

func pushThreeTest(s Stack, t *testing.T) {
	if got := s.Push(1); got != 1 {
		t.Errorf("Push(1) == %v, want %v", got, 1)
	}
	if got := s.Push(2); got != 2 {
		t.Errorf("Push(2) == %v, want %v", got, 2)
	}
	if got := s.Push(3); got != 3 {
		t.Errorf("Push(2) == %v, want %v", got, 3)
	}
	if got := s.Size(); got != 3 {
		t.Errorf("Size() == %v, want %v", got, 3)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != 3 || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, 3, nil)
	}

	if v, err := s.Pop(); v != 3 || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, 3, nil)
	}
	if got := s.Size(); got != 2 {
		t.Errorf("Size() == %v, want %v", got, 2)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != 2 || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, 2, nil)
	}

	if v, err := s.Pop(); v != 2 || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, 2, nil)
	}
	if got := s.Size(); got != 1 {
		t.Errorf("Size() == %v, want %v", got, 1)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != 1 || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, 1, nil)
	}

	if v, err := s.Pop(); v != 1 || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, 1, nil)
	}
	if got := s.Size(); got != 0 {
		t.Errorf("Size() == %v, want %v", got, 0)
	}
	if got := s.IsEmpty(); !got {
		t.Errorf("IsEmpty() == %v, want %v", got, true)
	}
	if _, err := s.Top(); err != ErrEmpty {
		t.Errorf("Top() should fail because stack is empty")
	}
	if _, err := s.Pop(); err != ErrEmpty {
		t.Errorf("Pop() should fail because stack is empty")
	}
}
