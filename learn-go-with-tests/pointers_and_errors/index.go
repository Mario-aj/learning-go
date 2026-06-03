package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

var InsufficientBalanceError = errors.New("It is not possible to withdraw: insufficient balance")

func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if quantity > w.balance {
		return InsufficientBalanceError
	}

	w.balance -= quantity
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
