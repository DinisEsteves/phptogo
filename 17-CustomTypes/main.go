package main

import "fmt"

func main() {
	type person struct {
		name     string
		location string
		age      int
	}

	person1 := person{
		name:     "Dinis",
		location: "Foo",
		age:      12,
	}

	fmt.Println(person1)
}
