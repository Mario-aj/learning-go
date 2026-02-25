package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			printLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
			os.Exit(-1)
		}
	}
}

func intro() {
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
	sites := readSitesFromFile()

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

	if res.StatusCode == http.StatusOK {
		fmt.Println("Site: ", site, "is up and running")
	} else {
		fmt.Println("Site: ", site, "is down. Status code: ", res.StatusCode)
	}

	logRegister(site, res.StatusCode == http.StatusOK)
}
func readSitesFromFile() []string {
	sites := make([]string, 0)

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		sites = append(sites, line)
	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error opening log file: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	fmt.Println("Showing logs...")

	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
