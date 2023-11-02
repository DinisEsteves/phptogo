package main

import "fmt"

func main() {
	person := map[string]map[string]string{
		"001": {
			"name": "Dinis",
			"age":  "25",
		},
	}

	fmt.Println(person["001"]["name"])
}
