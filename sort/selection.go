package sort

func SelectionSort(data []int, num int) {
	for i := 0; i < num; i++ {
		min := i
		for j := i + 1; j < num; j++ {
			if data[j] < data[min] {
				min = j
			}
		}

		if i != min {
			data[i], data[min] = data[min], data[i]
		}
	}

	return
}
