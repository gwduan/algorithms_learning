package queue

import (
	"sync"
	"testing"
)

func newQueueTest(q Queue, t *testing.T) {
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
	if _, err := q.Get(); err != ErrEmpty {
		t.Errorf("Get() should fail because queue is empty")
	}
}

func putOneTest(q Queue, t *testing.T) {
	if got, err := q.Put(1); got != 1 || err != nil {
		t.Errorf("Put(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
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

	if got, err := q.Get(); got != 1 || err != nil {
		t.Errorf("Get() == (%v, %v), want (%v, %v)", got, err, 1, nil)
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
	if _, err := q.Get(); err != ErrEmpty {
		t.Errorf("Get() should fail because queue is empty")
	}
}

func putTwoTest(q Queue, t *testing.T) {
	if got, err := q.Put(1); got != 1 || err != nil {
		t.Errorf("Put(1) == (%v, %v), want (%v, %v)", got, err, 1, nil)
	}
	if got, err := q.Put(2); got != 2 || err != nil {
		t.Errorf("Put(2) == (%v, %v), want (%v, %v)", got, err, 2, nil)
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

	if got, err := q.Get(); got != 1 || err != nil {
		t.Errorf("Get() == (%v, %v), want (%v, %v)", got, err, 1, nil)
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

	if got, err := q.Get(); got != 2 || err != nil {
		t.Errorf("Get() == (%v, %v), want (%v, %v)", got, err, 2, nil)
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
	if _, err := q.Get(); err != ErrEmpty {
		t.Errorf("Get() should fail because queue is empty")
	}
}

func lockTest(q Queue, t *testing.T) {
	chPut := make(chan int, 10)
	chGet := make(chan int, 10)
	chEmpty := make(chan int, 10)
	chFull := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var numPut, numFull int
			for j := 0; j < 10000; j++ {
				_, err := q.Put(j)
				switch err {
				case ErrFull:
					numFull++
				case nil:
					numPut++
				}
			}
			chPut <- numPut
			chFull <- numFull
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			var numGet, numEmpty int
			for j := 0; j < 10000; j++ {
				_, err := q.Get()
				switch err {
				case ErrEmpty:
					numEmpty++
				case nil:
					numGet++
				}
			}
			chGet <- numGet
			chEmpty <- numEmpty
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 10000; j++ {
				_ = q.Size()
				_ = q.IsEmpty()
				_, _ = q.Head()
				_, _ = q.Tail()
			}
		}()
	}
	wg.Wait()

	var totalPut, totalGet, totalFull, totalEmpty int
	close(chPut)
	close(chGet)
	close(chEmpty)
	close(chFull)
	for v := range chPut {
		totalPut += v
	}
	for v := range chGet {
		totalGet += v
	}
	for v := range chEmpty {
		totalEmpty += v
	}
	for v := range chFull {
		totalFull += v
	}
	if totalPut+totalFull != totalGet+totalEmpty {
		t.Errorf("Total put(%d)+full(%d) should == get(%d)+empty(%d)",
			totalPut, totalFull, totalGet, totalEmpty)
	}

	remainGet := 0
	for {
		if _, err := q.Get(); err != nil {
			break
		}
		remainGet++
	}
	if totalPut != totalGet+remainGet {
		t.Errorf("Put(%d) should == Get(%d)", totalPut,
			totalGet+remainGet)
	}
}
