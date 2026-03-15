package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
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

func (t *BST) Insert(value int) {
	insert(&t.root, value)
}

func (t *BST) Search(value int) bool {
	return search(t.root, value)
}

func main() {
	tree := BST{}

	fmt.Println(tree.root)

	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)

	fmt.Println(tree.root)
	fmt.Println(tree.root.left)
	fmt.Println(tree.root.right)

	fmt.Println("Search for number 4", tree.Search(4))
	fmt.Println("Search for number 0", tree.Search(0))
	fmt.Println("Search for number 9", tree.Search(9))
	fmt.Println("Search for number 2", tree.Search(2))
}
