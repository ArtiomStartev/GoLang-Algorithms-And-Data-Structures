package dijkstra

import (
	"math"
	"sort"
)

/**
 * Performs Dijkstra's algorithm to find the shortest path from a start node to all other nodes in the graph.
 * @param graph map[string]map[string]int - The graph represented as an adjacency list with edge weights
 * @param startNode string - The starting node
 * @returns map[string]int - An object containing the shortest distances from the start node to all other nodes
 */

func dijkstra(graph map[string]map[string]int, startNode string) map[string]int {
	var queue []string

	// Initialize distances map with all nodes set to the maximum integer value except the start node
	distances := make(map[string]int)
	for node, _ := range graph {
		queue = append(queue, node)

		if node == startNode {
			distances[node] = 0
		} else {
			distances[node] = math.MaxInt64
		}
	}

	// Sort the queue based on distances
	sort.Slice(queue, func(i, j int) bool {
		return distances[queue[i]] < distances[queue[j]]
	})

	// While the queue is not empty
	for len(queue) > 0 {
		// Dequeue the node from the queue
		currentNode := queue[0]
		queue = queue[1:]

		// Visit each neighbor of the current node
		for neighbor, distanceToNeighbor := range graph[currentNode] {
			totalDistance := distances[currentNode] + distanceToNeighbor

			// Update distance if a shorter path is found
			if totalDistance < distances[neighbor] {
				distances[neighbor] = totalDistance

				// Sort the queue again based on updated distances
				sort.Slice(queue, func(i, j int) bool {
					return distances[queue[i]] < distances[queue[j]]
				})
			}
		}
	}

	return distances
}
