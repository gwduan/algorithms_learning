package stack

import (
	"sync"
	"testing"
)

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
	if got, err := s.Push(1); err != nil || got != 1 {
		t.Errorf("Push(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
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
	if got, err := s.Push(1); err != nil || got != 1 {
		t.Errorf("Push(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := s.Push("h"); err != nil || got != "h" {
		t.Errorf("Push(\"h\") == (%v, %v), want (%v, %v)", got, err, "h", nil)
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

	if v, err := s.Pop(); v != "h" || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, "h", nil)
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
	if got, err := s.Push(1); err != nil || got != 1 {
		t.Errorf("Push(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := s.Push("h"); err != nil || got != "h" {
		t.Errorf("Push(\"h\") == (%v, %v), want (%v, %v)", got, err, "h", nil)
	}
	if got, err := s.Push(true); err != nil || got != true {
		t.Errorf("Push(3) == (%v, %v), want (%v, %v)", got, err, true, nil)
	}
	if got := s.Size(); got != 3 {
		t.Errorf("Size() == %v, want %v", got, 3)
	}
	if got := s.IsEmpty(); got {
		t.Errorf("IsEmpty() == %v, want %v", got, false)
	}
	if v, err := s.Top(); v != true || err != nil {
		t.Errorf("Top()= (%v, %v), want (%v, %v)", v, err, true, nil)
	}

	if v, err := s.Pop(); v != true || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, true, nil)
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

	if v, err := s.Pop(); v != "h" || err != nil {
		t.Errorf("Pop()= (%v, %v), want (%v, %v)", v, err, "h", nil)
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

func lockTest(s Stack, t *testing.T) {
	chPush := make(chan int, 10)
	chPop := make(chan int, 10)
	chEmpty := make(chan int, 10)
	chFull := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var numPush, numFull int
			for j := 0; j < 10000; j++ {
				_, err := s.Push(j)
				switch err {
				case ErrFull:
					numFull++
				case nil:
					numPush++
				}
			}
			chPush <- numPush
			chFull <- numFull
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			var numPop, numEmpty int
			for j := 0; j < 10000; j++ {
				_, err := s.Pop()
				switch err {
				case ErrEmpty:
					numEmpty++
				case nil:
					numPop++
				}
			}
			chPop <- numPop
			chEmpty <- numEmpty
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 10000; j++ {
				_ = s.Size()
				_ = s.IsEmpty()
				_, _ = s.Top()
			}
		}()
	}
	wg.Wait()

	var totalPush, totalPop, totalFull, totalEmpty int
	close(chPush)
	close(chPop)
	close(chEmpty)
	close(chFull)
	for v := range chPush {
		totalPush += v
	}
	for v := range chPop {
		totalPop += v
	}
	for v := range chEmpty {
		totalEmpty += v
	}
	for v := range chFull {
		totalFull += v
	}
	if totalPush+totalFull != totalPop+totalEmpty {
		t.Errorf("Total push(%d)+full(%d) should == pop(%d)+empty(%d)",
			totalPush, totalFull, totalPop, totalEmpty)
	}

	remainPop := 0
	for {
		if _, err := s.Pop(); err != nil {
			break
		}
		remainPop++
	}
	if totalPush != totalPop+remainPop {
		t.Errorf("Push(%d) should == Pop(%d)", totalPush,
			totalPop+remainPop)
	}
}
