package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	result := wallet.Balance()
	expected := 10

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
