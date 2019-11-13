package sort

func MergeSort(data []int, num int) {
	sort(data, 0, num-1)
}

func sort(data []int, low int, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2
	sort(data, low, mid)
	sort(data, mid+1, high)

	merge(data, low, mid, high)
}

func merge(data []int, low int, mid int, high int) {
	i := low
	j := mid + 1
	k := 0
	tmp := make([]int, high-low+1)
	for i <= mid && j <= high {
		if data[i] <= data[j] {
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
