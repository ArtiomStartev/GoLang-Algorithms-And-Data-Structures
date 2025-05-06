package recursive_tree_sum

type Node struct {
	Value    int
	Children []*Node
}

// RecursiveTreeSum calculates the sum of values in a tree structure using a recursive approach
func RecursiveTreeSum(tree []*Node) int {
	var sum int

	for _, node := range tree {
		sum += node.Value

		if len(node.Children) > 0 {
			sum += RecursiveTreeSum(node.Children)
		}
	}

	return sum
}
