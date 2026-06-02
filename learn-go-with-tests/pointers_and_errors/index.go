package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(quantity int) {
	w.balance += quantity
}

func (w *Wallet) Balance() int {
	return w.balance
}
