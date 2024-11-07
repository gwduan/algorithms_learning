package sort

import (
	"testing"
)

func TestInsertionSortInts(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 7}

	data1 := []int{2, 3, 5, 1, 7, 4}
	InsertionSort(data1, len(data1), cmpInts)
	if !sorted(data1, len(data1), cmpInts) {
		t.Errorf("got %v, want %v", data1, want)
	}

	data2 := []int{1, 2, 3, 4, 5, 7}
	InsertionSort(data2, len(data2), cmpInts)
	if !sorted(data2, len(data2), cmpInts) {
		t.Errorf("got %v, want %v", data2, want)
	}

	data3 := []int{7, 5, 4, 3, 2, 1}
	InsertionSort(data3, len(data3), cmpInts)
	if !sorted(data3, len(data3), cmpInts) {
		t.Errorf("got %v, want %v", data3, want)
	}
}

func TestInsertionSortFloats(t *testing.T) {
	want := []float64{1.2, 2.3, 3.4, 4.5, 5.6, 7.8}

	data1 := []float64{2.3, 3.4, 5.6, 1.2, 7.8, 4.5}
	InsertionSort(data1, len(data1), cmpFloats)
	if !sorted(data1, len(data1), cmpFloats) {
		t.Errorf("got %v, want %v", data1, want)
	}
}
