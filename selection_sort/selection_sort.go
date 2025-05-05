package selection_sort

// SelectionSort performs a selection sort on the given array of integers
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}

	return arr
}
