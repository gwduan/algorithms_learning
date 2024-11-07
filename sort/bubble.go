package sort

func BubbleSort[E any](data []E, num int, cmp func(E, E) int) {
	for i := 0; i < num; i++ {
		flag := false
		for j := 0; j < num-i-1; j++ {
			if cmp(data[j], data[j+1]) > 0 {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}
