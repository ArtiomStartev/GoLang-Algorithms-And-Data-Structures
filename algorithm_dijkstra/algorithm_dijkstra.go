package dijkstra

import (
	"math"
	"sort"
)

type Graph map[string]map[string]int
type Distances map[string]int

// Dijkstra finds the shortest path from the start node to all other nodes in the graph using Dijkstra's algorithm
func Dijkstra(graph Graph, startNode string) Distances {
	var queue []string
	distances := make(Distances)

	for node := range graph {
		queue = append(queue, node)

		if node == startNode {
			distances[node] = 0
		} else {
			distances[node] = math.MaxInt64
		}
	}

	sort.Slice(queue, func(i, j int) bool {
		return distances[queue[i]] < distances[queue[j]]
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

		sort.Slice(queue, func(i, j int) bool {
			return distances[queue[i]] < distances[queue[j]]
		})
	}

	return distances
}
