package encapsulation

import (
	"errors"
	"fmt"
)

type GoodBankAccount struct {
	balance float64
}

func NewBankAccount() *GoodBankAccount {
	return &GoodBankAccount{balance: 0.0}
}

func (g *GoodBankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount is lesser than the zero which is not allowed")
	}
	g.balance += amount
	return nil
}

func (g *GoodBankAccount) WithDraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount is lesser than zero which is not allowed")
	}
	if amount > g.balance {
		return errors.New("insufficient balance in your account")
	}
	g.balance -= amount
	return nil
}

func (g *GoodBankAccount) GetBalance() float64 {
	return g.balance
}

func GoodEncapsulation() {
	bankAccount := NewBankAccount()
	fmt.Println("Initial Balance", bankAccount.GetBalance())

	err := bankAccount.Deposit(45)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = bankAccount.WithDraw(40)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Final balance", bankAccount.GetBalance())
}
