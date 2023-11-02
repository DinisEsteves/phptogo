package main

import "fmt"

func main() {
	person := map[string]any{
		"name": "Dinis",
		"age":  "25",
	}

	// .(int) is not conversion it is assertion, is success ok = true
	age, ok := person["age"].(int)

	if ok {
		fmt.Println(age + 5)
	}
}
