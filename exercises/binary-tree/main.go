package main

import (
	"bst-in-go/tree"
	"fmt"
)

func main() {
	tree := tree.Bst{}

	fmt.Println("-------------- MIN --------------")
	minValue, err := tree.Min()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The min value  inside the BST is: ", minValue)
	}

	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(1)

	fmt.Println(tree.Root)
	fmt.Println(tree.Root.Left)
	fmt.Println(tree.Root.Right)

	fmt.Println("Search for number 4", tree.Search(4))
	fmt.Println("Search for number 0", tree.Search(0))
	fmt.Println("Search for number 9", tree.Search(9))
	fmt.Println("Search for number 2", tree.Search(2))

	// fmt.Println("-------------- MIN --------------")
	// minValue, err := tree.Min()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("The min value  inside the BST is: ", minValue)
	// }
}
