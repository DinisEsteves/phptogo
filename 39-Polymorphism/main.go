package main

import "fmt"

type account struct {
	balance int
}

type contactable interface {
	openingPhrase() string
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

func (a adultAccount) openingPhrase() string {
	return "This message is regarding your account."
}

type minorAccount struct {
	account
	limit int
}

func (a minorAccount) openingPhrase() string {
	return "This message is regarding your ward's account."
}

func (a *minorAccount) withdraw(ammout int) {

	if ammout > a.limit {
		fmt.Println("Account limit exceed")
	} else {
		a.account.withdraw(ammout)
	}
}

func sendMail(c contactable, message string) string {
	return c.openingPhrase() + " " + message
}

func main() {
	a := minorAccount{
		account: account{balance: 1000},
	}

	mail := sendMail(a, "Thanks!")

	fmt.Println(mail)
}
