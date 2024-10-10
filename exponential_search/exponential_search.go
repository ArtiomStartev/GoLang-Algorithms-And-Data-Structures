package exponential_search

import "algorithms/recursive_binary_search"

func exponentialSearch(arr []int, target int) int {
	// Return -1 if the array is empty
	if len(arr) == 0 {
		return -1
	}

	// Return the index of the target element if found
	if arr[0] == target {
		return 0
	}

	// Find the range for binary search by repeated doubling
	i := 1
	for i < len(arr) && arr[i] <= target {
		i *= 2
	}

	// Perform binary search in the range [i/2, min(i, len(arr))]
	return recursive_binary_search.RecursiveBinarySearch(arr, target, i/2, min(i, len(arr)))
}
