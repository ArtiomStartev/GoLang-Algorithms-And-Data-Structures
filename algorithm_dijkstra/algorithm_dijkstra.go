package dijkstra

import (
	"cmp"
	"math"
	"slices"
)

type Graph map[string]map[string]int
type Distances map[string]int

// Dijkstra finds the shortest path from the start node to all other nodes in the graph using Dijkstra's algorithm
func Dijkstra(graph Graph, startNode string) Distances {
	queue := make([]string, 0, len(graph))
	distances := make(Distances, len(graph))

	for node := range graph {
		queue = append(queue, node)

		if node == startNode {
			distances[node] = 0
		} else {
			distances[node] = math.MaxInt64
		}
	}

	slices.SortFunc(queue, func(a, b string) int {
		return cmp.Compare(distances[a], distances[b])
	})

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		for neighbor, distanceToNeighbor := range graph[currNode] {
			totalDistance := distances[currNode] + distanceToNeighbor

			if totalDistance < distances[neighbor] {
				distances[neighbor] = totalDistance
			}
		}

		slices.SortFunc(queue, func(a, b string) int {
			return cmp.Compare(distances[a], distances[b])
		})
	}

	return distances
}
