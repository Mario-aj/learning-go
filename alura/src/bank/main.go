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

	fmt.Println(account.balance)

	withdrawMessage := account.withdraw(100)
	fmt.Println(withdrawMessage)
	fmt.Println("Your new balance is: ", account.balance)

	fmt.Println("__________________________________")

	messaage, value := account.deposit(1250)
	fmt.Println(messaage)
	fmt.Printf("Your new balance is: %.2f \n", value)

	fmt.Println("_________________Transfer between accounts_______________")
	elmarAccount := Account{
		"Elmar Jorge",
		528,
		123456,
		0,
	}

	success := account.transfer(100, &elmarAccount)
	fmt.Println(success)

	fmt.Println(account.balance)
	fmt.Println(elmarAccount.balance)
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

func (account *Account) deposit(value float64) (string, float64) {
	if value <= 0 {
		return "You are trying to deposit a negative value", account.balance
	}

	account.balance += value
	return "Deposit successful", account.balance
}

func (account *Account) transfer(value float64, destinityAccount *Account) bool {
	if value > 0 && account.balance >= value {
		account.withdraw(value)
		destinityAccount.deposit(value)
		return true
	}

	return false
}
