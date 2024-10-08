package binary_search_tree

// Node represents a single node in the Binary Search Tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinarySearchTree represents the structure of a Binary Search Tree
type BinarySearchTree struct {
	Root *Node
}

// Search checks whether a value exists in the Binary Search Tree
// It returns true if the value is found, otherwise returns false
func (bst *BinarySearchTree) Search(value int) bool {
	return searchRecursive(bst.Root, value)
}

// searchRecursive is a helper function that recursively searches for a value in the Binary Search Tree
// It returns true if the value is found in the current node or its children, otherwise false
func searchRecursive(node *Node, value int) bool {
	if node == nil {
		return false // Value not found
	}

	if value == node.Value {
		return true // Value found
	} else if value < node.Value {
		return searchRecursive(node.Left, value) // Search in the left subtree
	} else {
		return searchRecursive(node.Right, value) // Search in the right subtree
	}
}
