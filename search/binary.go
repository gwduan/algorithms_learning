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
