package breadth_first_search

type GraphNode[T any] struct {
	Value     T
	Neighbors []*GraphNode[T]
}

// BFS performs a breadth-first search (BFS) on a graph represented by nodes
func BFS[T any](startNode *GraphNode[T]) []T {
	var traversalOrder []T
	var queue []*GraphNode[T]
	visited := make(map[*GraphNode[T]]bool)

	queue = append(queue, startNode)
	visited[startNode] = true

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		traversalOrder = append(traversalOrder, currNode.Value)

		for _, neighbor := range currNode.Neighbors {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return traversalOrder
}
