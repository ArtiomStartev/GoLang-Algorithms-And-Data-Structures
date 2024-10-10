package main

import (
	"algorithms/iterative_binary_search"
	"algorithms/linear_search"
	"algorithms/utils"
	"fmt"
	"math/rand"
)

func main() {
	slice := utils.GenerateSlice(100000000, utils.Ascending)
	target := slice[rand.Intn(len(slice))]

	// 1. Measure the execution time of the Linear Search algorithm
	utils.MeasureExecutionTime("Linear Search", func() {
		//fmt.Println("Slice to search in: ", slice)
		//fmt.Println("Target to search for: ", target)

		res := linear_search.LinearSearch(slice, target)
		if res != -1 {
			fmt.Printf("Target found at index: %d\n", res)
		} else {
			fmt.Println("Target not found in the slice.")
		}
	})

	// 2. Measure the execution time of the Binary Search algorithm
	utils.MeasureExecutionTime("Binary Search", func() {
		//fmt.Println("Slice to search in: ", slice)
		//fmt.Println("Target to search for: ", target)

		res := iterative_binary_search.IterativeBinarySearch(slice, target)
		if res == -1 {
			fmt.Println("Target not found in the slice.")
		} else {
			fmt.Printf("Target found at index: %d\n", res)
		}
	})
}
