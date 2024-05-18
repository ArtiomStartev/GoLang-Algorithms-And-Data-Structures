package main

/**
 * Performs Breadth-First Search on a graph starting from a specified vertex.
 * @param graph map[string][]string - The graph represented as an adjacency list
 * @param startVertex string - The vertex to start the search from
 * @returns []string - A slice representing the BFS traversal order
 */

func bfs(graph map[string][]string, startVertex string) []string {
	// Initialize an empty queue for BFS traversal
	var queue []string

	// Initialize a map to keep track of visited vertices
	visited := make(map[string]bool)

	// Initialize a slice to store the BFS traversal order
	var traversalOrder []string

	// Add the start vertex to the queue and mark it as visited
	queue = append(queue, startVertex)
	visited[startVertex] = true

	// Perform BFS traversal
	for len(queue) > 0 {
		// Dequeue a vertex from the queue
		currentVertex := queue[0]
		queue = queue[1:]

		// Add the current vertex to the traversal order
		traversalOrder = append(traversalOrder, currentVertex)

		// Get the neighbors of the current vertex from the graph
		neighbors := graph[currentVertex]

		// Iterate through the neighbors
		for _, neighbor := range neighbors {
			// If the neighbor has not been visited, mark it as visited and enqueue it
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return traversalOrder
}
