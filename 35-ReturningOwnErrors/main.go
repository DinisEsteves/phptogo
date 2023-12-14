package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInvalidEmail = errors.New("Invalid email exception")

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return ErrInvalidEmail
	}

	return nil
}

func main() {
	var email string

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	err := validateEmail(email)
	if err != nil {
		if err == ErrInvalidEmail {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Unexpected error")
		os.Exit(1)
	}

	fmt.Println("Email: ", email)
}
