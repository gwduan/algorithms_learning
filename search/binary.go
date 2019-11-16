package search

func BinarySearch(data []int, num int, value int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] == value {
			return mid
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchFirstEqualValue(data []int, num int, value int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > value {
			high = mid - 1
			continue
		}

		if data[mid] < value {
			low = mid + 1
			continue
		}

		if mid == 0 || data[mid-1] != value {
			return mid
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchLastEqualValue(data []int, num int, value int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > value {
			high = mid - 1
			continue
		}

		if data[mid] < value {
			low = mid + 1
			continue
		}

		if mid == num-1 || data[mid+1] != value {
			return mid
		} else {
			low = mid + 1
		}
	}

	return -1
}

func BinarySearchFirstGEValue(data []int, num int, value int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] < value {
			low = mid + 1
			continue
		}

		if mid == 0 || data[mid-1] < value {
			return mid
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchLastLEValue(data []int, num int, value int) int {
	low := 0
	high := num - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > value {
			high = mid - 1
			continue
		}

		if mid == num-1 || data[mid+1] > value {
			return mid
		} else {
			low = mid + 1
		}
	}

	return -1
}
