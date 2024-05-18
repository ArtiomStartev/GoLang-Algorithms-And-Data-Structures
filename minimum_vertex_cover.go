package main

import (
	"math"
)

/**
 * Finds a minimum vertex cover in a graph.
 * @param graph map[string][]string - The graph represented as an adjacency list.
 * @returns []string - A minimum vertex cover of the graph.
 */

func findMinimumVertexCover(graph map[string][]string) []string {
	// Initialize an empty slice for the minimum vertex cover
	var vertexCover []string

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
		vertexCover = append(vertexCover, minDegreeVertex)

		// Step 2: Remove the selected vertex and its neighborhood from the graph
		for _, neighbor := range graph[minDegreeVertex] {
			delete(graph, neighbor) // Remove the adjacent vertices
		}
		delete(graph, minDegreeVertex) // Remove the selected vertex
	}

	return vertexCover
}
