package main

import "testing"

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	}

	errorConfirmation := func(t *testing.T, err error, expected string) {
		t.Helper()

		if err == nil {
			t.Fatal("Expected an err, but no err occurred")
		}

		result := err.Error()

		if result != expected {
			t.Errorf("result %s, expected %s", result, expected)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		confirmBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		confirmBalance(t, wallet, initialBalance)
		errorConfirmation(t, err, "It is not possible to withdraw: insufficient balance")
	})
}
