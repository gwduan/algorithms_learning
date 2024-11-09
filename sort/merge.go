package sort

func MergeSort[E any](data []E, num int, cmp func(E, E) int) {
	mergeSort(data, 0, num-1, cmp)
}

func mergeSort[E any](data []E, low int, high int, cmp func(E, E) int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2
	mergeSort(data, low, mid, cmp)
	mergeSort(data, mid+1, high, cmp)

	merge(data, low, mid, high, cmp)
}

func merge[E any](data []E, low int, mid int, high int, cmp func(E, E) int) {
	i := low
	j := mid + 1
	k := 0
	tmp := make([]E, high-low+1)
	for i <= mid && j <= high {
		if cmp(data[i], data[j]) <= 0 {
			tmp[k] = data[i]
			i++
		} else {
			tmp[k] = data[j]
			j++
		}
		k++
	}

	s := i
	e := mid
	if j <= high {
		s = j
		e = high
	}
	for s <= e {
		tmp[k] = data[s]
		s++
		k++
	}

	for i := 0; i < k; i++ {
		data[low+i] = tmp[i]
	}
}
