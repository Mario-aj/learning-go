package main

import "fmt"


func main () {
	var numbers = [5]int{1, 2, 3, 4, 5}
	names := [3]string{"Mario", "Alfredo", "Jorge"}

	fmt.Println(numbers)
	fmt.Println(names)

	// length of the array
	fmt.Println("The Length  of the array of numbers is: ", len(numbers))
	fmt.Println("The length of the array of names is: ", len(names))

}
