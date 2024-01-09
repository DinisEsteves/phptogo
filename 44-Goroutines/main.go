package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello")
}

func main() {
	fmt.Println(time.Now().Second())

	wg := sync.WaitGroup{}
	wg.Add(2)

	go hello(&wg)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Bye")
	}()

	wg.Wait()

	fmt.Println(time.Now().Second())
}
