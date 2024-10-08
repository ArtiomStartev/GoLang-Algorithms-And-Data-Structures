package recursive_tree_sum

type Node struct {
	Value    int
	Children []Node
}

func recursiveTreeSum(tree []Node) int {
	var sum int

	for _, node := range tree {
		sum += node.Value

		if len(node.Children) > 0 {
			sum += recursiveTreeSum(node.Children)
		}
	}

	return sum
}
