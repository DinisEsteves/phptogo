package main

import "fmt"

type account struct {
	balance int
}

func (a *account) statement() {
	fmt.Println("Current balance:", a.balance)
}

func (a *account) withdraw(ammout int) {
	a.balance -= ammout
}

type adultAccount struct {
	account
}

type minorAccount struct {
	account
	limit int
}

func (a *minorAccount) withdraw(ammout int) {

	if ammout > a.limit {
		fmt.Println("Account limit exceed")
	} else {
		a.account.withdraw(ammout)
	}
}

func main() {
	a := minorAccount{
		account: account{balance: 1000},
	}

	a.withdraw(900)

	a.statement()
}
