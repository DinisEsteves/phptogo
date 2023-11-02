package main

import "fmt"

type sLevel int

func (sl sLevel) toString() string {
	switch sl {
	case 1:
		return "Junior"
	case 2:
		return "Senior"
	default:
		return "Undefined"
	}
}
func main() {
	var foo sLevel = 2

	fmt.Println(foo.toString())
}
