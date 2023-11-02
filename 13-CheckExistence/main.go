package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Dinis",
		"age":  "25",
	}

	profession, ok := person["profession"]

	if ok {
		fmt.Println("Profession set")
	} else {
		fmt.Println("Profession not set")
	}

	fmt.Println(profession)
}
