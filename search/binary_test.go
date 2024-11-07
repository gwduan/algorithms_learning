package search

import "testing"

func TestBinarySearch(t *testing.T) {
	data := []int{1, 3, 5, 6, 7, 14, 20}
	if got := BinarySearch(data, len(data), 7, cmpInts); got != 4 {
		t.Errorf("got %v, want %v", got, 4)
	}
}

func TestBinarySearchFirstEqualValue(t *testing.T) {
	data := []int{1, 3, 4, 5, 6, 8, 8, 8, 14, 20}
	if got := BinarySearchFirstEqualValue(data, len(data), 8, cmpInts); got != 5 {
		t.Errorf("got %v, want %v", got, 5)
	}
}

func TestBinarySearchLastEqualValue(t *testing.T) {
	data := []int{1, 3, 4, 5, 6, 7, 8, 8, 8, 20}
	if got := BinarySearchLastEqualValue(data, len(data), 8, cmpInts); got != 8 {
		t.Errorf("got %v, want %v", got, 8)
	}
}

func TestBinarySearchFirstGEValue(t *testing.T) {
	data := []int{1, 3, 4, 5, 6, 8, 8, 8, 14, 20}
	if got := BinarySearchFirstGEValue(data, len(data), 7, cmpInts); got != 5 {
		t.Errorf("got %v, want %v", got, 5)
	}
}

func TestBinarySearchLastLEValue(t *testing.T) {
	data := []int{1, 3, 4, 5, 6, 7, 7, 8, 9, 20}
	if got := BinarySearchLastLEValue(data, len(data), 7, cmpInts); got != 6 {
		t.Errorf("got %v, want %v", got, 6)
	}
}

func cmpInts(a, b int) int {
	return a - b
}
