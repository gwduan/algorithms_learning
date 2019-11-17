package skiplist

import "testing"

func TestSkipList(t *testing.T) {
	s := NewSkipList()
	if got := s.Find(3); got != nil {
		t.Errorf("find(3) = %v, want %v", got, nil)
	}

	s.Insert(1)
	s.Insert(5)
	s.Insert(3)
	if got := s.Find(3); got == nil || got.value != 3 {
		t.Errorf("find(3) = %v, want %v", got, 3)
	}
	if got := s.Find(6); got != nil {
		t.Errorf("find(6) = %v, want %v", got, nil)
	}

	s.Insert(2)
	s.Insert(8)
	s.Insert(7)
	s.Insert(9)
	s.Insert(4)
	s.Insert(6)
	if got := s.Find(8); got == nil || got.value != 8 {
		t.Errorf("find(8) = %v, want %v", got, 8)
	}
	if got := s.Find(10); got != nil {
		t.Errorf("find(10) = %v, want %v", got, nil)
	}

	s.Delete(16)
	if got := s.Find(8); got == nil || got.value != 8 {
		t.Errorf("find(8) = %v, want %v", got, 8)
	}
	s.Delete(6)
	if got := s.Find(6); got != nil {
		t.Errorf("find(6) = %v, want %v", got, nil)
	}
}
