package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		result := wallet.Balance()
		expected := Bitcoin(10)

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		result := wallet.Balance()
		expected := Bitcoin(10)

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	})
}
