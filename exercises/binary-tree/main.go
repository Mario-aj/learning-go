package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func createNode(value int) *Node {
	return &Node{value: value, left: nil, right: nil}
}

func insert(node **Node, value int) {
	if *node == nil {
		*node = createNode(value)
		return
	}

	if value > (*node).value {
		insert(&(*node).right, value)
	} else if value < (*node).value {
		insert(&(*node).left, value)
	}
}

func search(node *Node, value int) bool {
	if node == nil { 
		return false;
	}

	if node.value > value {
		return search(node.left, value)
	}

	if (node.value < value) {
		return search(node.right, value)
	}

	return true;
}

func main() {
	tree := BinaryTree{}

	fmt.Println(tree.root)

	insert(&tree.root, 5)
	insert(&tree.root, 8)
	insert(&tree.root, 4)
	insert(&tree.root, 7)
	insert(&tree.root, 2)

	fmt.Println(tree.root)
	fmt.Println(tree.root.left)
	fmt.Println(tree.root.right)

	fmt.Println("Search for number 4", search(tree.root, 4))
	fmt.Println("Search for number 0", search(tree.root, 0))
	fmt.Println("Search for number 9", search(tree.root, 9))
	fmt.Println("Search for number 2", search(tree.root, 2))
}
