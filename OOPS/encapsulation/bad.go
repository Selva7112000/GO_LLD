package encapsulation

import "fmt"

type BadBankAccount struct {
	Balance float64
}

func BadEncapsulation() {
	bankAccount := BadBankAccount{Balance: 0.0}
	bankAccount.Balance = -1
	fmt.Println(bankAccount.Balance)
}
