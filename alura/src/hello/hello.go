package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delay = 5 * time.Second
const numberOfChecks = 2

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
	sites := []string {"https://staging-app.planne.com.br", "https://www.google.com", "https://www.youtube.com"}

	for range numberOfChecks {
		for j, site := range sites {
			fmt.Println("Index: ", j, "\tSite: ", site)
			siteTester(site)
		}
		fmt.Printf("\n")
		time.Sleep(delay)
	}
}

func siteTester(site string) {
	res, _ := http.Get(site)

	if res.StatusCode == 200 {
		fmt.Println("Site: ", site, "is up and running")
	} else {
		fmt.Println("Site: ", site, "is down. Status code: ", res.StatusCode)
	}
}