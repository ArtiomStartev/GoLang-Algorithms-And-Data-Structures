package main

import (
	"math"
)

/**
 * Finds a maximal independent set in a graph.
 * @param graph map[string][]string - The graph represented as an adjacency list.
 * @returns []string - A slice representing the vertices of a maximal independent set.
 */

func findMaximalIndependentSet(graph map[string][]string) []string {
	// Initialize an empty slice for the independent set
	var independentSet []string

	for len(graph) > 0 {
		// Step 1: Select a vertex of minimum degree
		var minDegreeVertex string
		minDegree := math.MaxInt64 // Initialize to the maximum integer value

		for vertex, neighbors := range graph {
			if len(neighbors) < minDegree {
				minDegreeVertex = vertex
				minDegree = len(neighbors)
			}
		}

		// Append the selected vertex to the vertex cover slice
		independentSet = append(independentSet, minDegreeVertex)

		// Step 2: Remove the selected vertex and its neighborhood from the graph
		for _, neighbor := range graph[minDegreeVertex] {
			for _, vertex := range graph[neighbor] {
				graph[vertex] = filter(graph[vertex], neighbor)
			}
			delete(graph, neighbor) // Remove the adjacent vertices
		}
		delete(graph, minDegreeVertex) // Remove the selected vertex
	}

	return independentSet
}

// filter removes all instances of target from slice
func filter(slice []string, target string) []string {
	result := []string{}
	for _, item := range slice {
		if item != target {
			result = append(result, item)
		}
	}
	
	return result
}
