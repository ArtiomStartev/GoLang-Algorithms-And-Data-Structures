package main

/**
 * Linear search function to find the index of a target element in a slice.
 * @param slice []int - The slice to search in
 * @param target int - The target element to search for
 * @returns int - The index of the target element if found, otherwise -1
 */

func linearSearch(slice []int, target int) int {
	// Iterate through each element of the slice
	for i := 0; i < len(slice); i++ {
		// If the current element is equal to the target, return its index
		if slice[i] == target {
			return i
		}
	}

	// If the target element is not found, return -1
	return -1
}
