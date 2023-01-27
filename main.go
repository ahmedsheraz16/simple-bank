package main

import (
	"fmt"
)

type Account struct {
	accountName    string
	accountNumber  int
	accountBalance int
}

func (a *Account) Deposit(amount int) {
	a.accountBalance += amount
}
func (a *Account) Withdraw(amount int) {
	if a.accountBalance < amount {
		fmt.Print("Insufficient Balance")
	} else {
		a.accountBalance -= amount
	}

}
func (a *Account) CheckBalance() {
	fmt.Println("Balance", a.accountBalance)
}
func main() {
	acc := Account{}
	acc.accountNumber = 123
	//acc.accountBalance = acc.accountNumber

	acc.Deposit(100)
	acc.CheckBalance()

	acc.Withdraw(70)
	acc.CheckBalance()

	// acc.Withdraw(100)
	// acc.CheckBalance()
}
