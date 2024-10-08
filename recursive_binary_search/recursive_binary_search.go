package recursive_binary_search

/**
 * Recursively searches for a target element in a sorted slice using the Binary Search algorithm.
 * @param slice []int - The sorted slice to search in
 * @param target int - The target element to search for
 * @param start int - The start index of the current search range
 * @param end int - The end index of the current search range
 * @returns int - The index of the target element if found, otherwise -1
 */

func recursiveBinarySearch(slice []int, target, start, end int) int {
	// Calculate the middle index of the current search range
	middle := (start + end) / 2

	// Return the index of the target element if found
	if slice[middle] == target {
		return middle
	}

	// If the middle element is greater than the target, search in the left half
	if slice[middle] > target {
		return recursiveBinarySearch(slice, target, start, middle-1)
	}

	// If the middle element is less than the target, search in the right half
	if slice[middle] < target {
		return recursiveBinarySearch(slice, target, middle+1, end)
	}

	// Return -1 if the target element is not found
	return -1
}
