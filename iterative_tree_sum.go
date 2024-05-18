package main

/**
 * Performs an iterative depth-first traversal of a tree and computes the sum of all node values.
 * @param tree []Node - The tree structure represented as a slice of nodes.
 * @returns int - The sum of all node values in the tree.
 */

type Node struct {
	Value    int
	Children []Node
}

func iterativeTreeSum(tree []Node) int {
	var sum int

	stack := make([]Node, len(tree))
	copy(stack, tree)

	// Iterate through the stack
	for len(stack) > 0 {
		// Pop the top node from the stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Add the value of the node to the sum
		sum += node.Value

		// If the node has children, push them onto the stack
		if len(node.Children) > 0 {
			stack = append(stack, node.Children...)
		}
	}

	return sum
}
