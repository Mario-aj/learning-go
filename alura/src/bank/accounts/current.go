package accounts

type Current struct {
	Owner         string
	Agency        int
	AccountNumber int
	Balance       float64
}

func (account *Current) Withdraw(value float64) string {
	if value <= 0 {
		return "You are trying to withdraw a negative value"
	}

	if value > account.Balance {
		return "You do not have enough Balance"
	}

	account.Balance -= value
	return "Withdraw successful"
}

func (account *Current) Deposit(value float64) (string, float64) {
	if value <= 0 {
		return "You are trying to deposit a negative value", account.Balance
	}

	account.Balance += value
	return "Deposit successful", account.Balance
}

func (account *Current) Transfer(value float64, destinityAccount *Current) bool {
	if value > 0 && account.Balance >= value {
		account.Withdraw(value)
		destinityAccount.Deposit(value)
		return true
	}

	return false
}
