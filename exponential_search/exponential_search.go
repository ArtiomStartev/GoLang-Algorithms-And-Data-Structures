package exponential_search

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
	return binarySearch(arr, target, i/2, min(i, len(arr)))
}

func binarySearch(arr []int, target, start, end int) int {
	// Calculate the middle index of the current search range
	middle := (start + end) / 2

	// Return the index of the target element if found
	if arr[middle] == target {
		return middle
	}

	// If the middle element is greater than the target, search in the left half
	if arr[middle] > target {
		return binarySearch(arr, target, start, middle-1)
	}

	// If the middle element is less than the target, search in the right half
	if arr[middle] < target {
		return binarySearch(arr, target, middle+1, end)
	}

	// Return -1 if the target element is not found
	return -1
}
