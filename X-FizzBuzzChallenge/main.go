package main

import "fmt"

func main() {
	max := 100
	for i := 1; i <= max; i++ {
		if isMultiple(i, 5) && isMultiple(i, 3) {
			fmt.Println("Fizz")
		} else if isMultiple(i, 5) {
			fmt.Println("Buzz")
		} else if isMultiple(i, 3) {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}

func isMultiple(number int, divisor int) bool {
	return number%divisor == 0
}
