package main

/**
 * Sorts a slice in ascending order using the Quick Sort algorithm.
 * @param slice []int - The slice to be sorted
 * @returns []int - The sorted slice
 */

func quickSort(slice []int) []int {
	// Base case: If the slice has 0 or 1 element, it is already sorted
	if len(slice) <= 1 {
		return slice
	}

	// Choose a pivot index (middle element)
	pivotIndex := len(slice) / 2

	less := make([]int, 0)
	greater := make([]int, 0)

	for i := 0; i < len(slice); i++ {
		// Skip the pivot element
		if i == pivotIndex {
			continue
		}

		// If the element is less than the pivot, append it to the 'less' slice
		if slice[i] < slice[pivotIndex] {
			less = append(less, slice[i])
		}

		// If the element is greater than the pivot, append it to the 'greater' slice
		if slice[i] > slice[pivotIndex] {
			greater = append(greater, slice[i])
		}
	}

	// Recursively sort the 'less' and 'greater' slices and combine them with the pivot element
	return append(append(quickSort(less), slice[pivotIndex]), quickSort(greater)...)
}
