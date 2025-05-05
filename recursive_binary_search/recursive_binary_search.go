package recursive_binary_search

// RecursiveBinarySearch performs a recursive binary search on a sorted array of integers
func RecursiveBinarySearch(arr []int, x, low, high int) int {
	mid := low + (high-low)/2

	if arr[mid] == x {
		return mid
	}

	if arr[mid] > x {
		return RecursiveBinarySearch(arr, x, low, mid-1)
	}

	if arr[mid] < x {
		return RecursiveBinarySearch(arr, x, mid+1, high)
	}

	return -1
}
