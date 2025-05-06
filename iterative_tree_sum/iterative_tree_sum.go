package iterative_tree_sum

type Node struct {
	Value    int
	Children []*Node
}

// IterativeTreeSum calculates the sum of values in a tree structure using an iterative approach
func IterativeTreeSum(tree *Node) int {
	var sum int
	var stack []*Node

	stack = append(stack, tree)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sum += node.Value

		if len(node.Children) > 0 {
			stack = append(stack, node.Children...)
		}
	}

	return sum
}
