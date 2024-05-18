package main

/**
 * Sorts a slice in ascending order using the Selection Sort algorithm.
 * @param slice []int - The slice to be sorted
 * @returns []int The sorted slice
 */

func selectionSort(slice []int) []int {
	// Outer loop iterates through each element of the slice
	for i := 0; i < len(slice); i++ {
		minIndex := i // Initialize minIndex to the current index i

		// Inner loop iterates through the unsorted part of the slice
		for j := i + 1; j < len(slice); j++ {
			// Update minIndex if a smaller element is found
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}

		// If minIndex is different from i, swap the elements at indices i and minIndex
		if i != minIndex {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
		}
	}

	return slice // Return the sorted slice
}
