package main

import "fmt"

func main() {
	names := []string{"Xs", "Bar", "X"}

	for i, name := range names {
		fmt.Println("Hello ", name, " your positions is ", i)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println("Hello ", names[i])
	}

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
