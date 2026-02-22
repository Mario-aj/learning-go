package main

import "fmt"

func main() {
	name := "Mario"
	version := "1.0.0"

	fmt.Println("Hello mr. ", name)
	fmt.Println("Thist app is at version: ", version)

	fmt.Println("--------------------------------")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
	fmt.Println("--------------------------------")

	var command int
	fmt.Print("Enter a command: ")
	fmt.Scan(&command)

	fmt.Println("You chose ", command)
}