package quick_sort

// QuickSort performs a quick sort on the given array of integers
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)

	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

// partition partitions the array around a pivot element
func partition(arr []int) int {
	pivotIndex := len(arr) - 1
	i := -1

	for j := 0; j < pivotIndex; j++ {
		if arr[j] <= arr[pivotIndex] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[pivotIndex] = arr[pivotIndex], arr[i+1]

	return i + 1
}
