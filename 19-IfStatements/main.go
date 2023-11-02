package main

import "fmt"

func main() {
	names := []string{"foo", "bar"}
	//variables declared inside if's can only be assessed inside if scope
	if firstName := names[0]; firstName == "foo" {
		fmt.Println(firstName)
	}
}
