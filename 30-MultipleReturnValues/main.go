package main

import (
	"fmt"
	"phptogo/30-MultipleReturnValues/queue"
)

func main() {
	size, _ := queue.Enqueue("Foo Bar", 5)

	fmt.Println(size)
}
