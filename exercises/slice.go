package main

import "fmt"


func main () {
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println(fruits)

	fruits = append(fruits, "orange", "pineapple")
	fmt.Println(fruits)

	fruits = append(fruits, "mango", "peach")
	fmt.Println(fruits)

	for idx, fruit := range fruits {
		fmt.Println("Index: ", idx, "\tFruit: ", fruit)
	}

	fmt.Println("\nThe length of the slice is: ", len(fruits))

	fmt.Println("The capacity of the slice is: ", cap(fruits))
}
