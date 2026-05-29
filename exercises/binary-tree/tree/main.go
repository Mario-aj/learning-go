package tree

import "fmt"


type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Bst struct {
	Root *Node
}

func createNode(value int) *Node {
	return &Node{Value: value, Left: nil, Right: nil}
}

func insert(node **Node, value int) {
	if *node == nil {
		*node = createNode(value)
		return
	}

	if value > (*node).Value {
		insert(&(*node).Right, value)
	} else if value < (*node).Value {
		insert(&(*node).Left, value)
	}
}

func search(node *Node, value int) bool {
	if node == nil { 
		return false;
	}

	if node.Value > value {
		return search(node.Left, value)
	}

	if (node.Value < value) {
		return search(node.Right, value)
	}

	return true;
}

func min(node *Node) (int, error) {
	if node == nil {
		return 0, fmt.Errorf("There is no min value, because the BST is empty")
	}

	if node.Left == nil {
		return node.Value, nil
	}

	return min(node.Left)
}

func (t *Bst) Insert(value int) {
	insert(&t.Root, value)
}

func (t *Bst) Search(value int) bool {
	return search(t.Root, value)
}

func (t *Bst) Min() (int ,error) {
	return min(t.Root)
}
