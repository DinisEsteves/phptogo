package main

import "fmt"

func main() {
	names := []string{"Xs", "Bar", "X"}

	switch names[0] {
	case "Foo", "Bar":
		fmt.Println("Your name is foo or bar")
	case "X":
	default:
		fmt.Println("Name not found")
	}
}
