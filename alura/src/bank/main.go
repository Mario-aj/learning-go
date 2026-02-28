package main

import (
	"bank/accounts"
	"fmt"
)

func main() {
	account := accounts.Current{
		Owner:         "Mario Jorge",
		Agency:        528,
		AccountNumber: 123456,
		Balance:       559.99,
	}

	fmt.Println(account.Balance)

	withdrawMessage := account.Withdraw(100)
	fmt.Println(withdrawMessage)
	fmt.Println("Your new balance is: ", account.Balance)

	fmt.Println("__________________________________")

	messaage, value := account.Deposit(1250)
	fmt.Println(messaage)
	fmt.Printf("Your new balance is: %.2f \n", value)

	fmt.Println("_________________Transfer between accounts_______________")
	elmarAccount := accounts.Current{
		Owner:         "Elmar Jorge",
		Agency:        528,
		AccountNumber: 123456,
		Balance:       0,
	}

	success := account.Transfer(100, &elmarAccount)
	fmt.Println(success)

	fmt.Println(account.Balance)
	fmt.Println(elmarAccount.Balance)
}
