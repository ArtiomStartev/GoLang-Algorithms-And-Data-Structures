package independent_set

import (
	"math"
)

type Graph map[string][]string

// MaximumIndependentSet finds the maximum independent set of a graph using a greedy approach
func MaximumIndependentSet(graph Graph) []string {
	var independentSet []string

	for len(graph) > 0 {
		var minDegreeNode string
		minDegree := math.MaxInt64

		for node, neighbors := range graph {
			if len(neighbors) < minDegree {
				minDegreeNode = node
				minDegree = len(neighbors)
			}
		}

		independentSet = append(independentSet, minDegreeNode)

		for _, neighbor := range graph[minDegreeNode] {
			for _, neighborOfNeighbor := range graph[neighbor] {
				graph[neighborOfNeighbor] = filter(graph[neighborOfNeighbor], neighbor)
			}
			delete(graph, neighbor)
		}
		delete(graph, minDegreeNode)
	}

	return independentSet
}

// filter removes all occurrences of x from arr
func filter(arr []string, x string) []string {
	var result []string
	for _, val := range arr {
		if val != x {
			result = append(result, val)
		}
	}

	return result
}
