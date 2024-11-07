package sort

func InsertionSort[E any](data []E, num int, cmp func(E, E) int) {
	if num <= 1 {
		return
	}

	for i := 1; i < num; i++ {
		value := data[i]
		j := i - 1
		for ; j >= 0; j-- {
			if cmp(data[j], value) <= 0 {
				break
			}
			data[j+1] = data[j]
		}
		data[j+1] = value
	}
}
