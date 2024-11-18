package skiplist

import (
	"fmt"
	"testing"
)

func cmpInts(a, b any) int {
	return a.(int) - b.(int)
}

func TestNewSkipList(t *testing.T) {
	s := NewSkipList(cmpInts)
	if got := s.Level(); got != 0 {
		t.Errorf("Level() = %v, want %v", got, 0)
	}
	if got := s.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}
	if got, err := s.Find(3); err == nil {
		t.Errorf("Find(3) = (%v, %v), want (%v, %v)", got, err, nil, ErrNotFound)
	}
	if err := s.Delete(5); err == nil {
		t.Errorf("Delete(5) = %v, want %v", err, ErrNotFound)
	}
}

func TestInsertOneValue(t *testing.T) {
	s := NewSkipList(cmpInts)
	if err := s.Insert(5); err != nil {
		t.Errorf("Insert(5) = %v, want %v", err, nil)
	}
	if got := s.Level(); got != 1 {
		t.Errorf("Level() = %v, want %v", got, 1)
	}
	if got := s.Length(); got != 1 {
		t.Errorf("Length() = %v, want %v", got, 1)
	}
	if got, err := s.Find(3); err == nil {
		t.Errorf("Find(3) = (%v, %v), want (%v, %v)", got, err, nil, ErrNotFound)
	}
	if got, err := s.Find(5); got.Value() != 5 || err != nil {
		t.Errorf("Find(5) = (%v, %v), want (%v, %v)", got, err, 5, nil)
	}
	if err := s.Insert(5); err == nil {
		t.Errorf("Insert(5) = %v, want %v", err, ErrDuplicate)
	}
	if err := s.Delete(1); err == nil {
		t.Errorf("Delete(1) = %v, want %v", err, ErrNotFound)
	}
	if err := s.Delete(5); err != nil {
		t.Errorf("Delete(5) = %v, want %v", err, nil)
	}
	if got := s.Level(); got != 0 {
		t.Errorf("Level() = %v, want %v", got, 0)
	}
	if got := s.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}
}

func TestInsertMultiValues(t *testing.T) {
	s := NewSkipList(cmpInts)
	for i := 5; i <= 100; i += 5 {
		if err := s.Insert(i); err != nil {
			t.Errorf("Insert(%v) = %v, want %v", i, err, nil)
		}
	}
	if got := s.Length(); got != 20 {
		t.Errorf("Length() = %v, want %v", got, 20)
	}
	for i := 100; i > 0; i -= 4 {
		if err := s.Insert(i); i%(4*5) != 0 && err != nil {
			t.Errorf("Insert(%v) = %v, want %v", i, err, nil)
		} else if i%(4*5) == 0 && err == nil {
			t.Errorf("Insert(%v) = %v, want %v", i, err, ErrDuplicate)
		}
	}
	if got := s.Length(); got != 40 {
		t.Errorf("Length() = %v, want %v", got, 40)
	}
	if got, err := s.Find(40); got.Value() != 40 || err != nil {
		t.Errorf("Find(40) = (%v, %v), want (%v, %v)", got, err, 40, nil)
	}
	if got, err := s.Find(33); err == nil {
		t.Errorf("Find(33) = (%v, %v), want (%v, %v)", got, err, nil, ErrNotFound)
	}

	for i := 4; i <= 100; i += 4 {
		s.Delete(i)
	}
	if got := s.Length(); got != 15 {
		t.Errorf("Length() = %v, want %v", got, 15)
	}
	if got, err := s.Find(36); err == nil {
		t.Errorf("Find(36) = (%v, %v), want (%v, %v)", got, err, nil, ErrNotFound)
	}
	if got, err := s.Find(25); got.Value() != 25 || err != nil {
		t.Errorf("Find(25) = (%v, %v), want (%v, %v)", got, err, 25, nil)
	}
	for i := 5; i <= 100; i += 5 {
		s.Delete(i)
	}
	if got := s.Length(); got != 0 {
		t.Errorf("Length() = %v, want %v", got, 0)
	}
	if got := s.Level(); got != 0 {
		t.Errorf("Level() = %v, want %v", got, 0)
	}
}

func TestFindAll(t *testing.T) {
	s := NewSkipList(cmpInts)
	for i := 5; i <= 100; i += 5 {
		s.Insert(i)
	}
	for i := 100; i > 0; i -= 4 {
		s.Insert(i)
	}

	wantValues := make([]int, 0, s.Length())
	for i := 1; i <= 100; i++ {
		if i%4 == 0 || i%5 == 0 {
			wantValues = append(wantValues, i)
		}
	}

	gots, _ := s.FindAll()
	gotValues := make([]int, 0, len(gots))
	for _, v := range gots {
		gotValues = append(gotValues, v.Value().(int))
	}

	if len(wantValues) != len(gotValues) {
		t.Errorf("FindAll() numbers = %v, want %v", len(gotValues), len(wantValues))
	}
	for i, v := range gotValues {
		if v != wantValues[i] {
			t.Errorf("FindAll() no. %v = %v, want %v", i, v, wantValues[i])
		}
	}
}

func TestFindBetween(t *testing.T) {
	s := NewSkipList(cmpInts)
	for i := 5; i <= 100; i += 5 {
		s.Insert(i)
	}
	for i := 100; i > 0; i -= 4 {
		s.Insert(i)
	}

	wantValues := make([]int, 0, s.Length())
	for i := 64; i < 94; i++ {
		if i%4 == 0 || i%5 == 0 {
			wantValues = append(wantValues, i)
		}
	}

	gots, _ := s.FindBetween(64, 94)
	gotValues := make([]int, 0, len(gots))
	for _, v := range gots {
		gotValues = append(gotValues, v.Value().(int))
	}

	if len(wantValues) != len(gotValues) {
		t.Errorf("FindBetween(64, 94) numbers = %v, want %v", len(gotValues), len(wantValues))
	}
	for i, v := range gotValues {
		if v != wantValues[i] {
			t.Errorf("FindBetween(64, 94) no. %v = %v, want %v", i, v, wantValues[i])
		}
	}
}

func TestPrintInfo(t *testing.T) {
	s := NewSkipList(cmpInts)
	for i := 5; i <= 100; i += 5 {
		s.Insert(i)
	}
	for i := 100; i > 0; i -= 4 {
		s.Insert(i)
	}

	printAllValues(s)

	findValuePath(s, 4)
	findValuePath(s, 10)
	findValuePath(s, 12)
	findValuePath(s, 16)
	findValuePath(s, 48)
	findValuePath(s, 60)
	findValuePath(s, 76)
	findValuePath(s, 80)
	findValuePath(s, 100)

	findValuePath(s, 7)
	findValuePath(s, 27)
	findValuePath(s, 53)
	findValuePath(s, 77)
	findValuePath(s, 91)
}

func printAllValues(s *SkipList) {
	if s.length == 0 {
		return
	}

	fmt.Printf("SkipList content begin --------\n")
	fmt.Printf("level=%v, length=%v\n", s.level, s.length)
	for pre := s.head; pre.forwards[0] != nil; pre = pre.forwards[0] {
		fmt.Printf("(%v,%v) ", pre.forwards[0].value, pre.forwards[0].level)
	}
	fmt.Println()
	fmt.Printf("SkipList content end ----------\n")
}

func findValuePath(s *SkipList, value any) {
	times := 0
	fmt.Printf("Find %v path: ", value)
	pre := s.head
	for i := s.level - 1; i >= 0; i-- {
		for pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) < 0 {
			fmt.Printf("(%v,%v) -> ", pre.forwards[i].value, i+1)
			times++
			pre = pre.forwards[i]
		}

		if pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) > 0 {
			fmt.Printf("(%v,%v) -> ", pre.forwards[i].value, i+1)
			times++
		}

		if pre.forwards[i] != nil && s.cmp(pre.forwards[i].value, value) == 0 {
			times++
			fmt.Printf("(%v,%v), times=%v.\n", pre.forwards[i].value, i+1, times)
			return
		}
	}

	fmt.Printf("not found, times=%v.\n", times)
}
