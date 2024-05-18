package main

/**
 * Iterative function to implement Binary Search
 * @param slice []int - The sorted slice to search in
 * @param target int - The target value to search for
 * @returns int - The index of the target value if found, otherwise -1
 */

func iterativeBinarySearch(slice []int, target int) int {
	var start int
	end := len(slice) - 1

	for start <= end {
		// Find the middle index
		middle := (start + end) / 2

		// If target is found at middle, return its index
		if slice[middle] == target {
			return middle
		}

		// If target is less than the middle element, search in the left half
		if slice[middle] > target {
			end = middle - 1
		}

		// If target is greater than the middle element, search in the right half
		if slice[middle] < target {
			start = middle + 1
		}
	}

	// If target is not found, return -1
	return -1
}
