package exponential_search

import "algorithms/recursive_binary_search"

// ExponentialSearch performs an exponential search on a sorted array of integers
func ExponentialSearch(arr []int, x int) int {
	if arr[0] == x {
		return 0
	}

	i := 1
	for i < len(arr) && arr[i] <= x {
		if i >= len(arr)/2 { // Check for possible overflow
			i = len(arr)
			break
		}
		i *= 2
	}

	return recursive_binary_search.RecursiveBinarySearch(arr, x, i/2, i)
}
