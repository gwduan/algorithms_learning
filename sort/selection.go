package sort

func SelectionSort[E any](data []E, num int, cmp func(E, E) int) {
	for i := 0; i < num; i++ {
		min := i
		for j := i + 1; j < num; j++ {
			if cmp(data[j], data[min]) < 0 {
				min = j
			}
		}

		if i != min {
			data[i], data[min] = data[min], data[i]
		}
	}
}
