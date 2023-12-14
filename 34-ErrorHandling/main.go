package main

import (
	"fmt"
	"io"
	"os"
)

func cat(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	bytes, _ := io.ReadAll(file)

	return string(bytes), nil
}

func main() {
	content, err := cat("main.go")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Print(content)
}
