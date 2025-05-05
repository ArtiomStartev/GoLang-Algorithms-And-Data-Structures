package merge_sort

// MergeSort performs a merge sort on the given array of integers
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge is a helper function that merges two sorted arrays into one sorted array
func merge(left, right []int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	var i, j int

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	for i < len(left) {
		sorted = append(sorted, left[i])
		i++
	}

	for j < len(right) {
		sorted = append(sorted, right[j])
		j++
	}

	return sorted
}
