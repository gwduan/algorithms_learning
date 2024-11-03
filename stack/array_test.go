package stack

import "testing"

func TestNewFixedArrayStack(t *testing.T) {
	s := NewFixedArrayStack(2)
	newStackTest(s, t)
}

func TestNewdArrayStack(t *testing.T) {
	s := NewArrayStack(2)
	newStackTest(s, t)
}

func TestFixedArrayStackPushOne(t *testing.T) {
	s := NewFixedArrayStack(2)
	pushOneTest(s, t)
}

func TestArrayStackPushOne(t *testing.T) {
	s := NewArrayStack(2)
	pushOneTest(s, t)
}

func TestFixedArrayStackPushTwo(t *testing.T) {
	s := NewFixedArrayStack(2)
	pushTwoTest(s, t)
}

func TestArrayStackPushTwo(t *testing.T) {
	s := NewArrayStack(2)
	pushTwoTest(s, t)
}

func TestFixedArrayStackPushThree(t *testing.T) {
	s := NewFixedArrayStack(2)
	if got, err := s.Push(1); err != nil || got != 1 {
		t.Errorf("Push(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := s.Push("h"); err != nil || got != "h" {
		t.Errorf("Push(\"h\") == (%v, %v), want (%v, %v)", got, err, "h", nil)
	}
	if got, err := s.Push(true); err == nil {
		t.Errorf("Push(true) == (%v, %v), want (%v, %v)", got, err, true, ErrFull)
	}
	if got := s.Size(); got != 2 {
		t.Errorf("Size() == %v, want %v", got, 2)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != "h" || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, "h", nil)
	}
}

func TestArrayStackPushThree(t *testing.T) {
	s := NewArrayStack(2)
	pushThreeTest(s, t)
}

func TestFixedArrayStackLock(t *testing.T) {
	s := NewFixedArrayStack(1000)
	lockTest(s, t)
}

func TestArrayStackLock(t *testing.T) {
	s := NewArrayStack(100)
	lockTest(s, t)
}
