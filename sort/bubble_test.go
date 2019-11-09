package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 7}

	data1 := []int{2, 3, 5, 1, 7, 4}
	BubbleSort(data1, len(data1))
	if !sorted(data1, len(data1)) {
		t.Errorf("got %v, want %v", data1, want)
	}

	data2 := []int{1, 2, 3, 4, 5, 7}
	BubbleSort(data2, len(data2))
	if !sorted(data2, len(data2)) {
		t.Errorf("got %v, want %v", data2, want)
	}

	data3 := []int{7, 5, 4, 3, 2, 1}
	BubbleSort(data3, len(data3))
	if !sorted(data3, len(data3)) {
		t.Errorf("got %v, want %v", data3, want)
	}
}

func sorted(data []int, num int) bool {
	for i := num - 1; i > 0; i-- {
		if data[i] < data[i-1] {
			return false
		}
	}

	return true
}
