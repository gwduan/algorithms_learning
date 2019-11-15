package search

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []int{1, 3, 5, 6, 7, 14, 20}
	if got := BinarySearch(data, len(data), 7); got != 4 {
		t.Errorf("got %v, want %v", got, 4)
	}
}
