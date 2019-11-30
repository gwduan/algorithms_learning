package sort

func HeapSort(data []int, num int) {
	if num <= 1 {
		return
	}

	buildHeap(data, num)

	for k := num - 1; k > 0; k-- {
		tmp := data[0]
		data[0] = data[k]
		data[k] = tmp

		heapify(data, 0, k-1)
	}
}

func buildHeap(data []int, num int) {
	for i := (num - 1) / 2; i >= 0; i-- {
		heapify(data, i, num-1)
	}
}

func heapify(data []int, i int, n int) {
	for {
		maxPos := i
		if i*2+1 <= n && data[i] < data[i*2+1] {
			maxPos = i*2 + 1
		}
		if i*2+2 <= n && data[maxPos] < data[i*2+2] {
			maxPos = i*2 + 2
		}
		if maxPos == i {
			break
		}

		tmp := data[i]
		data[i] = data[maxPos]
		data[maxPos] = tmp

		i = maxPos
	}
}
