package main

type Tree struct {
	value       int
	left, right *Tree
}

func main() {

	var tree *Tree
	tree = addNode(tree, 3)
	tree = addNode(tree, 1)
	tree = addNode(tree, 2)
}

func addNode(tree *Tree, value int) *Tree {

	if tree == nil {
		tree = new(Tree)
		tree.value = value
		return tree
	}
	if value < tree.value {
		tree.left = addNode(tree.left, value)
	} else {
		tree.right = addNode(tree.right, value)
	}
	return tree
}
