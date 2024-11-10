package sort

func QuickSort[E any](data []E, num int, cmp func(E, E) int) {
	quickSort(data, 0, num-1, cmp)
}

func quickSort[E any](data []E, low int, high int, cmp func(E, E) int) {
	if low >= high {
		return
	}

	pivot := partition(data, low, high, cmp)

	quickSort(data, low, pivot-1, cmp)
	quickSort(data, pivot+1, high, cmp)
}

func partition[E any](data []E, low int, high int, cmp func(E, E) int) int {
	pivot := data[high]
	i := low
	for j := low; j < high; j++ {
		if cmp(data[j], pivot) < 0 {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	data[i], data[high] = data[high], data[i]

	return i
}
