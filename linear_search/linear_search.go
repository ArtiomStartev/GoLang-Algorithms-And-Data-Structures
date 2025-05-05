package linear_search

// LinearSearch performs a linear search on an array of integers
func LinearSearch(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}

	return -1
}
