package depth_first_search

type GraphNode[T any] struct {
	Value     T
	Neighbors []*GraphNode[T]
}

// DFS performs a depth-first search (DFS) on a graph represented by nodes
func DFS[T any](startNode *GraphNode[T]) []T {
	var traversalOrder []T
	visited := make(map[*GraphNode[T]]bool)

	return dfsTraversal[T](startNode, visited, traversalOrder)
}

// dfsTraversal recursively traverses the graph in depth-first order
func dfsTraversal[T any](node *GraphNode[T], visited map[*GraphNode[T]]bool, traversalOrder []T) []T {
	visited[node] = true
	traversalOrder = append(traversalOrder, node.Value)

	for _, neighbor := range node.Neighbors {
		if !visited[neighbor] {
			traversalOrder = dfsTraversal[T](neighbor, visited, traversalOrder)
		}
	}

	return traversalOrder
}
