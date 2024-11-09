package sort

func HeapSort[E any](data []E, num int, cmp func(E, E) int) {
	if num <= 1 {
		return
	}

	buildHeap(data, num, cmp)

	for k := num - 1; k > 0; k-- {
		data[0], data[k] = data[k], data[0]

		heapify(data, 0, k-1, cmp)
	}
}

func buildHeap[E any](data []E, num int, cmp func(E, E) int) {
	for i := (num - 1) / 2; i >= 0; i-- {
		heapify(data, i, num-1, cmp)
	}
}

func heapify[E any](data []E, i int, n int, cmp func(E, E) int) {
	for {
		maxPos := i
		if i*2+1 <= n && cmp(data[i], data[i*2+1]) < 0 {
			maxPos = i*2 + 1
		}
		if i*2+2 <= n && cmp(data[maxPos], data[i*2+2]) < 0 {
			maxPos = i*2 + 2
		}
		if maxPos == i {
			break
		}

		data[i], data[maxPos] = data[maxPos], data[i]

		i = maxPos
	}
}
