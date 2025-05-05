package binary_search_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

// Search performs a binary search on a Binary Search Tree
func (bst *BinarySearchTree) Search(x int) bool {
	return searchNode(bst.Root, x)
}

// searchNode is a helper function to search for a value in the tree
func searchNode(node *Node, x int) bool {
	if node == nil {
		return false
	}

	if x == node.Value {
		return true
	} else if x < node.Value {
		return searchNode(node.Left, x)
	} else {
		return searchNode(node.Right, x)
	}
}
