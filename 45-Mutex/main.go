package main

import (
	"fmt"
	"sync"
)

var languages []string
var lock sync.Mutex

func addLanguage(lang string, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	languages = append(languages, lang)
	lock.Unlock()
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	go addLanguage("PHP", &wg)
	go addLanguage("GO", &wg)
	go addLanguage("JS", &wg)
	go addLanguage("RUBY", &wg)
	go addLanguage("TS", &wg)
	wg.Wait()

	fmt.Println(languages)
}
