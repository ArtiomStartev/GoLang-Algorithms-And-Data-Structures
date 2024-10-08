package bubble_sort

/**
 * Sorts a slice in ascending order using the Bubble Sort algorithm.
 * @param slice []int - The slice to be sorted
 * @returns []int - The sorted slice
 */

func bubbleSort(slice []int) []int {
	// Outer loop iterates through each element of the slice
	for i := 0; i < len(slice); i++ {
		// Inner loop iterates through the unsorted portion of the slice
		for j := 0; j < len(slice)-i-1; j++ {
			// Compare adjacent elements and swap them if they are in the wrong order
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j] // Swap the elements
			}
		}
	}

	return slice // Return the sorted slice
}
