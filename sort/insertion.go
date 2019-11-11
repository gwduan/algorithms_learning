package sort

func InsertionSort(data []int, num int) {
	if num <= 1 {
		return
	}

	for i := 1; i < num; i++ {
		value := data[i]
		j := i - 1
		for ; j >= 0; j-- {
			if data[j] > value {
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = value
	}

	return
}
