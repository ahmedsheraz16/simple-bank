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
func (a *Account) Transfer(amount int, b *Account) {
	if a.accountBalance < amount {
		fmt.Println("Insufficient funds")
	} else {
		a.accountBalance -= amount
		b.accountBalance += amount
		fmt.Printf("%v transferred amount from account %d to account %d", amount, a.accountNumber, b.accountNumber)
	}
}

func main() {
	account1 := Account{accountNumber: 1}
	account2 := Account{accountNumber: 2}
	//acc := make([]*Account, 2)
	account1.Deposit(400)
	fmt.Println("account1:")
	account1.CheckBalance()

	account2.Deposit(700)
	fmt.Println("account2:")
	account2.CheckBalance()
	// for _, ac := range acc {
	// 	ac.accountNumber = 123
	// 	//acc.accountBalance = acc.accountNumber

	// ac.CheckBalance()

	// ac.Withdraw(75)
	// ac.CheckBalance()
	// ac.accountNumber++

	// acc.Withdraw(100)
	// acc.CheckBalance()
	account1.Transfer(400, &account2)
	fmt.Println("Account1:")
	account1.CheckBalance()
	fmt.Println("Account2:")
	account2.CheckBalance()
}
