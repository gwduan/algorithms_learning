package sort

func cmpInts(a, b int) int {
	return a - b
}

func cmpFloats(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func sorted[E any](data []E, num int, cmp func(E, E) int) bool {
	for i := num - 1; i > 0; i-- {
		if cmp(data[i], data[i-1]) < 0 {
			return false
		}
	}

	return true
}
