package main

import "fmt"

type Account struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func main() {
	account := Account{
		"Mario Jorge",
		528,
		123456,
		559.99,
	}

	fmt.Println(account)
}
