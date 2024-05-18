package main

type TreeNode struct {
	Value    int
	Children []TreeNode
}

func recursiveTreeSum(tree []TreeNode) int {
	var sum int

	for _, node := range tree {
		sum += node.Value

		if len(node.Children) > 0 {
			sum += recursiveTreeSum(node.Children)
		}
	}

	return sum
}
