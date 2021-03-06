package sort

import "testing"

type sortFunc func([]int, int)

func TestBubbleSort(t *testing.T) {
	sortFuncTest(BubbleSort, t)
}

func TestInsertionSort(t *testing.T) {
	sortFuncTest(InsertionSort, t)
}

func TestSelectionSort(t *testing.T) {
	sortFuncTest(SelectionSort, t)
}

func TestMergeSort(t *testing.T) {
	sortFuncTest(MergeSort, t)
}

func TestQuickSort(t *testing.T) {
	sortFuncTest(QuickSort, t)
}

func TestHeapSort(t *testing.T) {
	sortFuncTest(HeapSort, t)
}

func sortFuncTest(sf sortFunc, t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 7}

	data1 := []int{2, 3, 5, 1, 7, 4}
	sf(data1, len(data1))
	if !sorted(data1, len(data1)) {
		t.Errorf("got %v, want %v", data1, want)
	}

	data2 := []int{1, 2, 3, 4, 5, 7}
	sf(data2, len(data2))
	if !sorted(data2, len(data2)) {
		t.Errorf("got %v, want %v", data2, want)
	}

	data3 := []int{7, 5, 4, 3, 2, 1}
	sf(data3, len(data3))
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
