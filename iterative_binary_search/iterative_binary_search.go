package iterative_binary_search

// IterativeBinarySearch performs an iterative binary search on a sorted array of integers
func IterativeBinarySearch(arr []int, x int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == x {
			return x
		}

		if arr[mid] > x {
			high = mid - 1
		}

		if arr[mid] < x {
			low = mid + 1
		}
	}

	return -1
}
