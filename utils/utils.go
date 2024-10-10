package utils

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"time"
)

type OrderType string

const (
	Ascending  OrderType = "ASC"
	Descending OrderType = "DESC"
	Random     OrderType = "RAND"
)

// MeasureExecutionTime takes a function and its name as arguments, executes the function,
// and prints the time taken for its execution.
func MeasureExecutionTime(name string, fn func()) {
	startTime := time.Now() // Record the start time

	fn() // Execute the function

	duration := time.Since(startTime) // Calculate the time difference
	fmt.Printf("Execution time of %s: %v\n", name, duration)
}

// GenerateSlice generates a slice of the specified size filled with random numbers based on the order type.
func GenerateSlice(size int, orderType OrderType) []int {
	slice := make([]int, size)

	// Fill the slice with random numbers
	for i := range slice {
		slice[i] = rand.Intn(math.MaxInt16)
	}

	// Sort the slice based on the order type
	switch orderType {
	case Ascending:
		slices.Sort(slice)
	case Descending:
		slices.SortFunc(slice, func(i, j int) int { return j - i })
	case Random:
		// Already filled with random numbers
	}

	return Unique(slice)
}

// Unique returns a new slice containing only the unique elements of the input slice.
func Unique(slice []int) []int {
	var result []int           // The resulting slice
	seen := make(map[int]bool) // A map to keep track of the unique elements

	// Iterate over the input slice
	for _, val := range slice {
		// If the element is not in the map, add it to the result slice and mark it as seen
		if ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}

	// Return the resulting slice
	return result
}
