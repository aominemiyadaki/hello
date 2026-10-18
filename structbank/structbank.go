package structbank

import "errors"

type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(deposit float64) {
	b.balance += deposit
}

func (b BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) Withdraw(value float64) error {

	if value > b.balance {
		return errors.New("ошбика списания, сумма списания больше баланса")
	}

	b.balance -= value
	return nil
}
