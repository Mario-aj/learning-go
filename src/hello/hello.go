package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	intro()

	for {
		showMenu()

		command := readInput()

		switch command {
			case 1: 
				startMonitoring()
			case 2:
				fmt.Println("Showing logs...")
			case 0:
				fmt.Println("Exiting...")
				os.Exit(0)
			default:
				fmt.Println("Invalid command")
				os.Exit(-1)
		}
	}
}

func intro () {
	name := "Mario"
	version := "1.0.0"

	fmt.Println("Hello mr. ", name)
	fmt.Println("Thist app is at version: ", version)
}

func showMenu() {
	fmt.Println("--------------------------------")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
	fmt.Println("--------------------------------")
}

func readInput() int {
	var command int
	fmt.Print("Enter a command: ")
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://httpbin.org/status/200"

	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Printf("Site %s is up and running\n", site)
	} else {
		fmt.Printf("Site %s is down. Status code: %d\n", site, res.StatusCode)
	}
}