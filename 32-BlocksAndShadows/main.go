package main

import "fmt"

func main() {
	name := "foo"
	func() {
		//shadowing
		name := "bar"
		fmt.Println(name)
	}()
	fmt.Println(name)

}
