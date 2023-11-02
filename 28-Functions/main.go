package main

import (
	"fmt"
	"phptogo/28-Functions/queue"
)

func main() {
	size := queue.EnqueueWithoutTimeout("Foo Bar")

	fmt.Println(size)
}
