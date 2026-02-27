package main

import "fmt"

type Account struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func (account *Account) withdraw(value float64) string {
	if value <= 0 {
		return "You are trying to withdraw a negative value"
	}

	if value > account.balance {
		return "You do not have enough balance"
	}

	account.balance -= value
	return "Withdraw successful"
}

func main() {
	account := Account{
		"Mario Jorge",
		528,
		123456,
		559.99,
	}

	fmt.Println(account.balance)

	withdrawMessage := account.withdraw(100)
	fmt.Println(withdrawMessage)
	fmt.Println("Your new balance is: ", account.balance)
}
