package depth_first_search

/**
 * Performs Depth-First Search on a graph starting from a specified vertex.
 * @param graph map[string][]string - The graph represented as an adjacency list
 * @param vertex string - The vertex to start the search from
 * @param traversalOrder []string - Slice to store the DFS traversal order
 * @param visited map[string]bool - Map to keep track of visited vertices
 */

func dfs(graph map[string][]string, vertex string, traversalOrder []string, visited map[string]bool) []string {
	traversalOrder = append(traversalOrder, vertex) // Add the current vertex to the traversal order
	visited[vertex] = true                          // Mark the current vertex as visited

	// Get the neighbors of the current vertex from the graph
	neighbors := graph[vertex]

	// Visit all adjacent vertices that have not been visited
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			traversalOrder = dfs(graph, neighbor, traversalOrder, visited)
		}
	}

	// Return the complete traversal order
	return traversalOrder
}

// Function to perform DFS traversal on the graph
func dfsTraversal(graph map[string][]string, startVertex string) []string {
	// Initialize a slice to store the DFS traversal order
	var traversalOrder []string

	// Initialize a map to keep track of visited vertices
	visited := make(map[string]bool)

	// Perform DFS traversal from the start vertex
	return dfs(graph, startVertex, traversalOrder, visited)
}
