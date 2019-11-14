package sort

func QuickSort(data []int, num int) {
	quickSort(data, 0, num-1)
}

func quickSort(data []int, low int, high int) {
	if low >= high {
		return
	}

	pivot := partition(data, low, high)

	quickSort(data, low, pivot-1)
	quickSort(data, pivot+1, high)
}

func partition(data []int, low int, high int) int {
	pivot := data[high]
	i := low
	for j := low; j < high; j++ {
		if data[j] < pivot {
			tmp := data[i]
			data[i] = data[j]
			data[j] = tmp
			i++
		}
	}

	tmp := data[i]
	data[i] = data[high]
	data[high] = tmp

	return i
}
