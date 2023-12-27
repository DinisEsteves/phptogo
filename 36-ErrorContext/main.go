package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrInvalidEmail = errors.New("Invalid email exception")

type InvalidEmailError struct {
	message   string
	idRequest string
}

func (err InvalidEmailError) Error() string {
	return err.message
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return InvalidEmailError{message: "invalid email", idRequest: "134"}
	}

	return nil
}

func main() {
	var email string

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	err := validateEmail(email)
	if err != nil {
		ieError, ok := err.(InvalidEmailError)
		if ok {
			fmt.Println(ieError.message)
			fmt.Println(ieError.idRequest)
			os.Exit(1)
		}
		fmt.Println("Unexpected error")
		os.Exit(1)
	}

	fmt.Println("Email: ", email)
}
