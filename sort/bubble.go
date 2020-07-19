package sort

func BubbleSort(data []int, num int) {
	for i := 0; i < num; i++ {
		flag := false
		for j := 0; j < num-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}
