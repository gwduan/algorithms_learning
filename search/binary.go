package search

func BinarySearch[E any](data []E, num int, value E, cmp func(E, E) int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2

		result := cmp(data[mid], value)
		if result == 0 {
			return mid
		} else if result < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchFirstEqualValue[E any](data []E, num int, value E, cmp func(E, E) int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2

		result := cmp(data[mid], value)
		if result > 0 {
			high = mid - 1
			continue
		}
		if result < 0 {
			low = mid + 1
			continue
		}

		if mid == 0 || cmp(data[mid-1], value) != 0 {
			return mid
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchLastEqualValue[E any](data []E, num int, value E, cmp func(E, E) int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2

		result := cmp(data[mid], value)
		if result > 0 {
			high = mid - 1
			continue
		}
		if result < 0 {
			low = mid + 1
			continue
		}

		if mid == num-1 || cmp(data[mid+1], value) != 0 {
			return mid
		} else {
			low = mid + 1
		}
	}

	return -1
}

func BinarySearchFirstGEValue[E any](data []E, num int, value E, cmp func(E, E) int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2

		if cmp(data[mid], value) < 0 {
			low = mid + 1
			continue
		}

		if mid == 0 || cmp(data[mid-1], value) < 0 {
			return mid
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchLastLEValue[E any](data []E, num int, value E, cmp func(E, E) int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2

		if cmp(data[mid], value) > 0 {
			high = mid - 1
			continue
		}

		if mid == num-1 || cmp(data[mid+1], value) > 0 {
			return mid
		} else {
			low = mid + 1
		}
	}

	return -1
}
